#!/usr/bin/env python
"""Simple tool for generating a client library.

Relevant links:
  https://developers.google.com/discovery/v1/reference/apis#resource
"""

import datetime

from six.moves import urllib_parse

from apitools.base.py import base_cli
from apitools.gen import command_registry
from apitools.gen import message_registry
from apitools.gen import service_registry
from apitools.gen import util


def _StandardQueryParametersSchema(discovery_doc):
    """Sets up dict of standard query parameters."""
    standard_query_schema = {
        'id': 'StandardQueryParameters',
        'type': 'object',
        'description': 'Query parameters accepted by all methods.',
        'properties': discovery_doc.get('parameters', {}),
    }
    # We add an entry for the trace, since Discovery doesn't.
    standard_query_schema['properties']['trace'] = {
        'type': 'string',
        'description': base_cli.TRACE_HELP,
        'location': 'query',
    }
    return standard_query_schema


def _ComputePaths(package, version, discovery_doc):
    full_path = urllib_parse.urljoin(
        discovery_doc['rootUrl'], discovery_doc['servicePath'])
    api_path_component = '/'.join((package, version, ''))
    if api_path_component not in full_path:
        return full_path, ''
    prefix, _, suffix = full_path.rpartition(api_path_component)
    return prefix + api_path_component, suffix


class DescriptorGenerator(object):

    """Code generator for a given discovery document."""

    def __init__(self, discovery_doc, client_info, names, root_package, outdir,
                 base_package, generate_cli=False, use_proto2=False,
                 unelidable_request_methods=None):
        self.__discovery_doc = discovery_doc
        self.__client_info = client_info
        self.__outdir = outdir
        self.__use_proto2 = use_proto2
        self.__description = util.CleanDescription(
            self.__discovery_doc.get('description', ''))
        self.__package = self.__client_info.package
        self.__version = self.__client_info.version
        self.__revision = discovery_doc.get('revision', '1')
        self.__generate_cli = generate_cli
        self.__root_package = root_package
        self.__base_files_package = base_package
        self.__base_files_target = (
            '//cloud/bigscience/apitools/base/py:apitools_base')
        self.__names = names
        self.__base_url, self.__base_path = _ComputePaths(
            self.__package, self.__client_info.url_version,
            self.__discovery_doc)

        # Order is important here: we need the schemas before we can
        # define the services.
        self.__message_registry = message_registry.MessageRegistry(
            self.__client_info, self.__names, self.__description,
            self.__root_package, self.__base_files_package)
        schemas = self.__discovery_doc.get('schemas', {})
        for schema_name, schema in schemas.items():
            self.__message_registry.AddDescriptorFromSchema(
                schema_name, schema)

        # We need to add one more message type for the global parameters.
        standard_query_schema = _StandardQueryParametersSchema(
            self.__discovery_doc)
        self.__message_registry.AddDescriptorFromSchema(
            standard_query_schema['id'], standard_query_schema)

        # Now that we know all the messages, we need to correct some
        # fields from MessageFields to EnumFields.
        self.__message_registry.FixupMessageFields()

        self.__command_registry = command_registry.CommandRegistry(
            self.__package, self.__version, self.__client_info,
            self.__message_registry, self.__root_package,
            self.__base_files_package, self.__base_url, self.__names)
        self.__command_registry.AddGlobalParameters(
            self.__message_registry.LookupDescriptorOrDie(
                'StandardQueryParameters'))

        self.__services_registry = service_registry.ServiceRegistry(
            self.__client_info,
            self.__message_registry,
            self.__command_registry,
            self.__base_url,
            self.__base_path,
            self.__names,
            self.__root_package,
            self.__base_files_package,
            unelidable_request_methods or [])
        services = self.__discovery_doc.get('resources', {})
        for service_name, methods in sorted(services.items()):
            self.__services_registry.AddServiceFromResource(
                service_name, methods)
        # We might also have top-level methods.
        api_methods = self.__discovery_doc.get('methods', [])
        if api_methods:
            self.__services_registry.AddServiceFromResource(
                'api', {'methods': api_methods})
        self.__client_info = self.__client_info._replace(
            scopes=self.__services_registry.scopes)

    @property
    def client_info(self):
        return self.__client_info

    @property
    def discovery_doc(self):
        return self.__discovery_doc

    @property
    def names(self):
        return self.__names

    @property
    def outdir(self):
        return self.__outdir

    @property
    def package(self):
        return self.__package

    @property
    def use_proto2(self):
        return self.__use_proto2

    def _GetPrinter(self, out):
        printer = util.SimplePrettyPrinter(out)
        return printer

    def WriteInit(self, out):
        """Write a simple __init__.py for the generated client."""
        printer = self._GetPrinter(out)
        printer('"""Common imports for generated %s client library."""',
                self.__client_info.package)
        printer('# pylint:disable=wildcard-import')
        printer()
        printer('import pkgutil')
        printer()
        printer('from %s import *', self.__base_files_package)
        if self.__root_package == '.':
            import_prefix = ''
        else:
            import_prefix = '%s.' % self.__root_package
        if self.__generate_cli:
            printer('from %s%s import *',
                    import_prefix, self.__client_info.cli_rule_name)
        printer('from %s%s import *',
                import_prefix, self.__client_info.client_rule_name)
        printer('from %s%s import *',
                import_prefix, self.__client_info.messages_rule_name)
        printer()
        printer('__path__ = pkgutil.extend_path(__path__, __name__)')

    def WriteIntermediateInit(self, out):
        """Write a simple __init__.py for an intermediate directory."""
        printer = self._GetPrinter(out)
        printer('#!/usr/bin/env python')
        printer('"""Shared __init__.py for apitools."""')
        printer()
        printer('from pkgutil import extend_path')
        printer('__path__ = extend_path(__path__, __name__)')

    def WriteSetupPy(self, out):
        """Write a setup.py for upload to PyPI."""
        printer = self._GetPrinter(out)
        year = datetime.datetime.now().year
        printer('# Copyright %s Google Inc. All Rights Reserved.' % year)
        printer('#')
        printer('# Licensed under the Apache License, Version 2.0 (the'
                '"License");')
        printer('# you may not use this file except in compliance with '
                'the License.')
        printer('# You may obtain a copy of the License at')
        printer('#')
        printer('#   http://www.apache.org/licenses/LICENSE-2.0')
        printer('#')
        printer('# Unless required by applicable law or agreed to in writing, '
                'software')
        printer('# distributed under the License is distributed on an "AS IS" '
                'BASIS,')
        printer('# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either '
                'express or implied.')
        printer('# See the License for the specific language governing '
                'permissions and')
        printer('# limitations under the License.')
        printer()
        printer('import setuptools')
        printer('REQUIREMENTS = [')
        with printer.Indent(indent='    '):
            # TODO(craigcitro): Have this track apitools' version.
            printer('"google-apitools>=0.4.8",')
            printer('"httplib2>=0.9",')
            printer('"oauth2client>=1.4.12",')
            printer('"protorpc>=0.10.0",')
        printer(']')
        printer('_PACKAGE = "apitools.clients.%s"' % self.__package)
        printer()
        printer('setuptools.setup(')
        # TODO(craigcitro): Allow customization of these options.
        with printer.Indent(indent='    '):
            printer('name="google-apitools-%s-%s",',
                    self.__package, self.__version)
            printer('version="0.4.%s",', self.__revision)
            printer('description="Autogenerated apitools library for %s",' % (
                self.__package,))
            printer('url="https://github.com/google/apitools",')
            printer('author="Craig Citro",')
            printer('author_email="craigcitro@google.com",')
            printer('packages=setuptools.find_packages(),')
            printer('install_requires=REQUIREMENTS,')
            printer('classifiers=[')
            with printer.Indent(indent='    '):
                printer('"Programming Language :: Python :: 2.7",')
                printer('"License :: OSI Approved :: Apache Software '
                        'License",')
            printer('],')
            printer('license="Apache 2.0",')
            printer('keywords="apitools apitools-%s %s",' % (
                self.__package, self.__package))
        printer(')')

    def WriteMessagesFile(self, out):
        self.__message_registry.WriteFile(self._GetPrinter(out))

    def WriteMessagesProtoFile(self, out):
        self.__message_registry.WriteProtoFile(self._GetPrinter(out))

    def WriteServicesProtoFile(self, out):
        self.__services_registry.WriteProtoFile(self._GetPrinter(out))

    def WriteClientLibrary(self, out):
        self.__services_registry.WriteFile(self._GetPrinter(out))

    def WriteCli(self, out):
        self.__command_registry.WriteFile(self._GetPrinter(out))
