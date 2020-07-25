// Code generated by protoc-gen-go.
// source: cq.proto
// DO NOT EDIT!

/*
Package cq is a generated protocol buffer package.

It is generated from these files:
	cq.proto

It has these top-level messages:
	Config
	Rietveld
	Gerrit
	Verifiers
*/
package cq

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This message describes a Commit Queue configuration. The config file cq.cfg
// should be stored in the config directory located on the branch that this CQ
// should commit to.
type Config struct {
	// Required. Version of the config format.
	Version *int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Required. Name of the CQ. May only contain characters [a-zA-Z0-9_]. It is
	// used for various purposes, including, but not limited to match the project
	// name for CLs on Rietveld, name of the project in the status app, internal
	// name for logging etc. CQ name should not be confused with the project name
	// in LUCI as there may be multiple CQs per project.
	CqName *string `protobuf:"bytes,2,opt,name=cq_name" json:"cq_name,omitempty"`
	// List of verifiers that verify if the CL is ready to be committed.
	Verifiers *Verifiers `protobuf:"bytes,3,opt,name=verifiers" json:"verifiers,omitempty"`
	// URL of the CQ status app to push updates to.
	CqStatusUrl *string `protobuf:"bytes,4,opt,name=cq_status_url" json:"cq_status_url,omitempty"`
	// When true, hash of the commit is not posted by CQ. This is used for
	// projects using gnumbd as latter publishes actual hash later. Default value
	// is false.
	HideRefInCommittedMsg *bool `protobuf:"varint,5,opt,name=hide_ref_in_committed_msg" json:"hide_ref_in_committed_msg,omitempty"`
	// Delay between commit bursts in seconds. Default value is 480.
	CommitBurstDelay *int32 `protobuf:"varint,6,opt,name=commit_burst_delay" json:"commit_burst_delay,omitempty"`
	// Maximum number of commits done sequentially, before waiting for
	// commit_burst_delay. Default value is 4.
	MaxCommitBurst *int32 `protobuf:"varint,7,opt,name=max_commit_burst" json:"max_commit_burst,omitempty"`
	// Defines whether a CQ is used in production. Allows to disable CQ for a
	// given branch. Default is true.
	InProduction *bool `protobuf:"varint,8,opt,name=in_production" json:"in_production,omitempty"`
	// Configuration options for Rietveld code review.
	Rietveld *Rietveld `protobuf:"bytes,9,opt,name=rietveld" json:"rietveld,omitempty"`
	// EXPERIMENTAL! Configuration options for Gerrit code review.
	// TODO(tandrii): update this doc (GERRIT).
	Gerrit *Gerrit `protobuf:"bytes,15,opt,name=gerrit" json:"gerrit,omitempty"`
	// This can be used to override the Git repository URL used to checkout and
	// commit changes on CQ host. This should only be used in case, when the
	// source repository is not supported by luci-config (e.g. GitHub).
	GitRepoUrl *string `protobuf:"bytes,10,opt,name=git_repo_url" json:"git_repo_url,omitempty"`
	// Target ref to commit to. This can be used to specify a different ref than
	// the one where the luci config is located. This is useful, e.g. for projects
	// that use gnumbd where CQ should commit into a pending ref.
	TargetRef *string `protobuf:"bytes,11,opt,name=target_ref" json:"target_ref,omitempty"`
	// Deprecated. URL of the SVN repository. We are deprecating SVN support.
	SvnRepoUrl *string `protobuf:"bytes,12,opt,name=svn_repo_url" json:"svn_repo_url,omitempty"`
	// If present, the CQ will refrain from processing any commits whose start
	// time is >= this time.
	//
	// This is an UTC RFC3339 (stiptime(tm)) string representing the time.
	DrainingStartTime *string `protobuf:"bytes,13,opt,name=draining_start_time" json:"draining_start_time,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Config) GetCqName() string {
	if m != nil && m.CqName != nil {
		return *m.CqName
	}
	return ""
}

func (m *Config) GetVerifiers() *Verifiers {
	if m != nil {
		return m.Verifiers
	}
	return nil
}

func (m *Config) GetCqStatusUrl() string {
	if m != nil && m.CqStatusUrl != nil {
		return *m.CqStatusUrl
	}
	return ""
}

func (m *Config) GetHideRefInCommittedMsg() bool {
	if m != nil && m.HideRefInCommittedMsg != nil {
		return *m.HideRefInCommittedMsg
	}
	return false
}

func (m *Config) GetCommitBurstDelay() int32 {
	if m != nil && m.CommitBurstDelay != nil {
		return *m.CommitBurstDelay
	}
	return 0
}

func (m *Config) GetMaxCommitBurst() int32 {
	if m != nil && m.MaxCommitBurst != nil {
		return *m.MaxCommitBurst
	}
	return 0
}

func (m *Config) GetInProduction() bool {
	if m != nil && m.InProduction != nil {
		return *m.InProduction
	}
	return false
}

func (m *Config) GetRietveld() *Rietveld {
	if m != nil {
		return m.Rietveld
	}
	return nil
}

func (m *Config) GetGerrit() *Gerrit {
	if m != nil {
		return m.Gerrit
	}
	return nil
}

func (m *Config) GetGitRepoUrl() string {
	if m != nil && m.GitRepoUrl != nil {
		return *m.GitRepoUrl
	}
	return ""
}

func (m *Config) GetTargetRef() string {
	if m != nil && m.TargetRef != nil {
		return *m.TargetRef
	}
	return ""
}

func (m *Config) GetSvnRepoUrl() string {
	if m != nil && m.SvnRepoUrl != nil {
		return *m.SvnRepoUrl
	}
	return ""
}

func (m *Config) GetDrainingStartTime() string {
	if m != nil && m.DrainingStartTime != nil {
		return *m.DrainingStartTime
	}
	return ""
}

type Rietveld struct {
	// Required. URL of the codereview site.
	Url *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// List of regular expressions used to check if CL's base URL should be
	// processed by this CQ. This may be useful if a single branch has multiple
	// sub-directories that are handled by different CQs. When no regular
	// expressions are specified, the regular expression '.*', which matches any
	// directory, is used.
	ProjectBases     []string `protobuf:"bytes,2,rep,name=project_bases" json:"project_bases,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Rietveld) Reset()                    { *m = Rietveld{} }
func (m *Rietveld) String() string            { return proto.CompactTextString(m) }
func (*Rietveld) ProtoMessage()               {}
func (*Rietveld) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rietveld) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Rietveld) GetProjectBases() []string {
	if m != nil {
		return m.ProjectBases
	}
	return nil
}

// Gerrit CQ is EXPERIMENTAL! See http://crbug.com/493899 for more info.
//
// Unlike Rietveld, Gerrit doesn't need a separate url.
// Instead, the git_repo_url must be specified on the Gerrit instance,
// and CQ will deduce Gerrit url from it.
//
// TODO(tandrii): support Rietveld and Gerrit at the same time.
// This basically requires to start two CQ instances, instead of one.
//
// For example, if https://chromium.googlesource.com/infra/infra.git is your
// repo url provided in `git_repo_url` above, then
// https://chromium-review.googlesource.com/#/admin/projects/infra/infra should
// show general properties of your project.
//
// Also,
// https://chromium-review.googlesource.com/#/admin/projects/infra/infra,access
// should show ACLs for refs in your project, but you may need to be admin to
// see it. This will come handy to enable and customize the CQ-related workflows
// for your project.
type Gerrit struct {
	// Optional. If set, tells CQ vote on a given label to mark result of CQ run.
	// The vote is either -1 if failed or 1 on success, and will be given on
	// non-dry runs only.
	// This vote can then be used in Gerrit's rule for submitting issues, so as to
	// require CQ run. CQ will attempt to submit issue only after setting this
	// label.
	CqVerifiedLabel  *string `protobuf:"bytes,1,opt,name=cq_verified_label" json:"cq_verified_label,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Gerrit) Reset()                    { *m = Gerrit{} }
func (m *Gerrit) String() string            { return proto.CompactTextString(m) }
func (*Gerrit) ProtoMessage()               {}
func (*Gerrit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Gerrit) GetCqVerifiedLabel() string {
	if m != nil && m.CqVerifiedLabel != nil {
		return *m.CqVerifiedLabel
	}
	return ""
}

// Verifiers are various types of checks that a Commit Queue performs on a CL.
// All verifiers must pass in order for a CL to be landed. Configuration file
// describes types of verifiers that should be applied to each CL and their
// parameters.
type Verifiers struct {
	// This verifier is used to ensure that an LGTM was posted to the code review
	// site from a valid project committer.
	// This verifier is not supported with Gerrit.
	ReviewerLgtm *Verifiers_ReviewerLgtmVerifier `protobuf:"bytes,1,opt,name=reviewer_lgtm" json:"reviewer_lgtm,omitempty"`
	// This verifier is used to check tree status before committing a CL. If the
	// tree is closed, then the verifier will wait until it is reopened.
	TreeStatus *Verifiers_TreeStatusLgtmVerifier `protobuf:"bytes,2,opt,name=tree_status" json:"tree_status,omitempty"`
	// This verifier triggers a set of tryjobs that are to be run on builders on
	// Buildbot. It automatically retries failed try-jobs and only allows CL to
	// land if each builder has succeeded in the latest retry. If a given tryjob
	// result is too old (>1 day) it is ignored.
	TryJob *Verifiers_TryJobVerifier `protobuf:"bytes,3,opt,name=try_job" json:"try_job,omitempty"`
	// This verifier is used to ensure that the author has signed Google's
	// Contributor License Agreement.
	SignCla          *Verifiers_SignCLAVerifier `protobuf:"bytes,4,opt,name=sign_cla" json:"sign_cla,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Verifiers) Reset()                    { *m = Verifiers{} }
func (m *Verifiers) String() string            { return proto.CompactTextString(m) }
func (*Verifiers) ProtoMessage()               {}
func (*Verifiers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Verifiers) GetReviewerLgtm() *Verifiers_ReviewerLgtmVerifier {
	if m != nil {
		return m.ReviewerLgtm
	}
	return nil
}

func (m *Verifiers) GetTreeStatus() *Verifiers_TreeStatusLgtmVerifier {
	if m != nil {
		return m.TreeStatus
	}
	return nil
}

func (m *Verifiers) GetTryJob() *Verifiers_TryJobVerifier {
	if m != nil {
		return m.TryJob
	}
	return nil
}

func (m *Verifiers) GetSignCla() *Verifiers_SignCLAVerifier {
	if m != nil {
		return m.SignCla
	}
	return nil
}

type Verifiers_ReviewerLgtmVerifier struct {
	// Required. Name of the chrome-infra-auth group, which contains the list of
	// identities authorized to approve (lgtm) a CL and trigger CQ run or dry
	// run.
	CommitterList *string `protobuf:"bytes,1,opt,name=committer_list" json:"committer_list,omitempty"`
	// Number of seconds to wait for LGTM on CQ. Default value is 0.
	MaxWaitSecs *int32 `protobuf:"varint,2,opt,name=max_wait_secs" json:"max_wait_secs,omitempty"`
	// Message to be posted to code review site when no LGTM is found. Default
	// value is "No LGTM from a valid reviewer yet. Only full committers are "
	// "accepted.\nEven if an LGTM may have been provided, it was from a "
	// "non-committer,\n_not_ a full super star committer.\nSee "
	// "http://www.chromium.org/getting-involved/become-a-committer\nNote that "
	// "this has nothing to do with OWNERS files."
	NoLgtmMsg *string `protobuf:"bytes,3,opt,name=no_lgtm_msg" json:"no_lgtm_msg,omitempty"`
	// Optional, but recommended. Name of the chrome-infra-auth group,
	// which contains the list of identities authorized to trigger CQ dry run.
	// This is usually the same group as tryjob-access.
	DryRunAccessList *string `protobuf:"bytes,4,opt,name=dry_run_access_list" json:"dry_run_access_list,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Verifiers_ReviewerLgtmVerifier) Reset()         { *m = Verifiers_ReviewerLgtmVerifier{} }
func (m *Verifiers_ReviewerLgtmVerifier) String() string { return proto.CompactTextString(m) }
func (*Verifiers_ReviewerLgtmVerifier) ProtoMessage()    {}
func (*Verifiers_ReviewerLgtmVerifier) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *Verifiers_ReviewerLgtmVerifier) GetCommitterList() string {
	if m != nil && m.CommitterList != nil {
		return *m.CommitterList
	}
	return ""
}

func (m *Verifiers_ReviewerLgtmVerifier) GetMaxWaitSecs() int32 {
	if m != nil && m.MaxWaitSecs != nil {
		return *m.MaxWaitSecs
	}
	return 0
}

func (m *Verifiers_ReviewerLgtmVerifier) GetNoLgtmMsg() string {
	if m != nil && m.NoLgtmMsg != nil {
		return *m.NoLgtmMsg
	}
	return ""
}

func (m *Verifiers_ReviewerLgtmVerifier) GetDryRunAccessList() string {
	if m != nil && m.DryRunAccessList != nil {
		return *m.DryRunAccessList
	}
	return ""
}

type Verifiers_TreeStatusLgtmVerifier struct {
	// Required. URL of the project tree status app.
	TreeStatusUrl    *string `protobuf:"bytes,1,opt,name=tree_status_url" json:"tree_status_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Verifiers_TreeStatusLgtmVerifier) Reset()         { *m = Verifiers_TreeStatusLgtmVerifier{} }
func (m *Verifiers_TreeStatusLgtmVerifier) String() string { return proto.CompactTextString(m) }
func (*Verifiers_TreeStatusLgtmVerifier) ProtoMessage()    {}
func (*Verifiers_TreeStatusLgtmVerifier) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 1}
}

func (m *Verifiers_TreeStatusLgtmVerifier) GetTreeStatusUrl() string {
	if m != nil && m.TreeStatusUrl != nil {
		return *m.TreeStatusUrl
	}
	return ""
}

type Verifiers_TryJobVerifier struct {
	// Buckets on which tryjobs are triggered/watched.
	Buckets []*Verifiers_TryJobVerifier_Bucket `protobuf:"bytes,1,rep,name=buckets" json:"buckets,omitempty"`
	// Provides project specific trybot retry configuration. This overrides the
	// defaults used in the CQ.
	TryJobRetryConfig *Verifiers_TryJobVerifier_TryJobRetryConfig `protobuf:"bytes,2,opt,name=try_job_retry_config" json:"try_job_retry_config,omitempty"`
	XXX_unrecognized  []byte                                      `json:"-"`
}

func (m *Verifiers_TryJobVerifier) Reset()                    { *m = Verifiers_TryJobVerifier{} }
func (m *Verifiers_TryJobVerifier) String() string            { return proto.CompactTextString(m) }
func (*Verifiers_TryJobVerifier) ProtoMessage()               {}
func (*Verifiers_TryJobVerifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 2} }

func (m *Verifiers_TryJobVerifier) GetBuckets() []*Verifiers_TryJobVerifier_Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *Verifiers_TryJobVerifier) GetTryJobRetryConfig() *Verifiers_TryJobVerifier_TryJobRetryConfig {
	if m != nil {
		return m.TryJobRetryConfig
	}
	return nil
}

type Verifiers_TryJobVerifier_Builder struct {
	// Name of the builder.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Optionally specify a builder name that triggers the given builder.
	// Otherwise, CQ will trigger this builder (default). If in doubt, you
	// probably won't need this.
	TriggeredBy *string `protobuf:"bytes,2,opt,name=triggered_by" json:"triggered_by,omitempty"`
	// When this field is present, it marks given builder as experimental. It
	// is only executed on a given percentage of the CLs and the outcome does
	// not affect the decicion whether a CL can land or not. This is typically
	// used to test new builders and estimate their capacity requirements.
	ExperimentPercentage *float32 `protobuf:"fixed32,4,opt,name=experiment_percentage" json:"experiment_percentage,omitempty"`
	XXX_unrecognized     []byte   `json:"-"`
}

func (m *Verifiers_TryJobVerifier_Builder) Reset()         { *m = Verifiers_TryJobVerifier_Builder{} }
func (m *Verifiers_TryJobVerifier_Builder) String() string { return proto.CompactTextString(m) }
func (*Verifiers_TryJobVerifier_Builder) ProtoMessage()    {}
func (*Verifiers_TryJobVerifier_Builder) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 2, 0}
}

func (m *Verifiers_TryJobVerifier_Builder) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Verifiers_TryJobVerifier_Builder) GetTriggeredBy() string {
	if m != nil && m.TriggeredBy != nil {
		return *m.TriggeredBy
	}
	return ""
}

func (m *Verifiers_TryJobVerifier_Builder) GetExperimentPercentage() float32 {
	if m != nil && m.ExperimentPercentage != nil {
		return *m.ExperimentPercentage
	}
	return 0
}

type Verifiers_TryJobVerifier_Bucket struct {
	// Name of the bucket. This is typically the same as a main name without
	// the 'main.' prefix, e.g. 'chromium.linux' or 'tryserver.webrtc'. CQ
	// will automatically add 'main.' prefix if not there.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Builders on which tryjobs should be triggered.
	Builders         []*Verifiers_TryJobVerifier_Builder `protobuf:"bytes,2,rep,name=builders" json:"builders,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *Verifiers_TryJobVerifier_Bucket) Reset()         { *m = Verifiers_TryJobVerifier_Bucket{} }
func (m *Verifiers_TryJobVerifier_Bucket) String() string { return proto.CompactTextString(m) }
func (*Verifiers_TryJobVerifier_Bucket) ProtoMessage()    {}
func (*Verifiers_TryJobVerifier_Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 2, 1}
}

func (m *Verifiers_TryJobVerifier_Bucket) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Verifiers_TryJobVerifier_Bucket) GetBuilders() []*Verifiers_TryJobVerifier_Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

type Verifiers_TryJobVerifier_TryJobRetryConfig struct {
	// Retry quota for a single tryjob.
	TryJobRetryQuota *int32 `protobuf:"varint,1,opt,name=try_job_retry_quota" json:"try_job_retry_quota,omitempty"`
	// Retry quota for all tryjobs in a CL.
	GlobalRetryQuota *int32 `protobuf:"varint,2,opt,name=global_retry_quota" json:"global_retry_quota,omitempty"`
	// The weight assigned to each tryjob failure.
	FailureRetryWeight *int32 `protobuf:"varint,3,opt,name=failure_retry_weight" json:"failure_retry_weight,omitempty"`
	// The weight assigned to each transient failure.
	TransientFailureRetryWeight *int32 `protobuf:"varint,4,opt,name=transient_failure_retry_weight" json:"transient_failure_retry_weight,omitempty"`
	// The weight assigned to tryjob timeouts.
	TimeoutRetryWeight *int32 `protobuf:"varint,5,opt,name=timeout_retry_weight" json:"timeout_retry_weight,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *Verifiers_TryJobVerifier_TryJobRetryConfig) Reset() {
	*m = Verifiers_TryJobVerifier_TryJobRetryConfig{}
}
func (m *Verifiers_TryJobVerifier_TryJobRetryConfig) String() string {
	return proto.CompactTextString(m)
}
func (*Verifiers_TryJobVerifier_TryJobRetryConfig) ProtoMessage() {}
func (*Verifiers_TryJobVerifier_TryJobRetryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 2, 2}
}

func (m *Verifiers_TryJobVerifier_TryJobRetryConfig) GetTryJobRetryQuota() int32 {
	if m != nil && m.TryJobRetryQuota != nil {
		return *m.TryJobRetryQuota
	}
	return 0
}

func (m *Verifiers_TryJobVerifier_TryJobRetryConfig) GetGlobalRetryQuota() int32 {
	if m != nil && m.GlobalRetryQuota != nil {
		return *m.GlobalRetryQuota
	}
	return 0
}

func (m *Verifiers_TryJobVerifier_TryJobRetryConfig) GetFailureRetryWeight() int32 {
	if m != nil && m.FailureRetryWeight != nil {
		return *m.FailureRetryWeight
	}
	return 0
}

func (m *Verifiers_TryJobVerifier_TryJobRetryConfig) GetTransientFailureRetryWeight() int32 {
	if m != nil && m.TransientFailureRetryWeight != nil {
		return *m.TransientFailureRetryWeight
	}
	return 0
}

func (m *Verifiers_TryJobVerifier_TryJobRetryConfig) GetTimeoutRetryWeight() int32 {
	if m != nil && m.TimeoutRetryWeight != nil {
		return *m.TimeoutRetryWeight
	}
	return 0
}

type Verifiers_SignCLAVerifier struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Verifiers_SignCLAVerifier) Reset()                    { *m = Verifiers_SignCLAVerifier{} }
func (m *Verifiers_SignCLAVerifier) String() string            { return proto.CompactTextString(m) }
func (*Verifiers_SignCLAVerifier) ProtoMessage()               {}
func (*Verifiers_SignCLAVerifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 3} }

func init() {
	proto.RegisterType((*Config)(nil), "Config")
	proto.RegisterType((*Rietveld)(nil), "Rietveld")
	proto.RegisterType((*Gerrit)(nil), "Gerrit")
	proto.RegisterType((*Verifiers)(nil), "Verifiers")
	proto.RegisterType((*Verifiers_ReviewerLgtmVerifier)(nil), "Verifiers.ReviewerLgtmVerifier")
	proto.RegisterType((*Verifiers_TreeStatusLgtmVerifier)(nil), "Verifiers.TreeStatusLgtmVerifier")
	proto.RegisterType((*Verifiers_TryJobVerifier)(nil), "Verifiers.TryJobVerifier")
	proto.RegisterType((*Verifiers_TryJobVerifier_Builder)(nil), "Verifiers.TryJobVerifier.Builder")
	proto.RegisterType((*Verifiers_TryJobVerifier_Bucket)(nil), "Verifiers.TryJobVerifier.Bucket")
	proto.RegisterType((*Verifiers_TryJobVerifier_TryJobRetryConfig)(nil), "Verifiers.TryJobVerifier.TryJobRetryConfig")
	proto.RegisterType((*Verifiers_SignCLAVerifier)(nil), "Verifiers.SignCLAVerifier")
}

var fileDescriptor0 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0x55, 0x08, 0x49, 0x9c, 0x31, 0x90, 0x0f, 0x13, 0xc0, 0x98, 0xaf, 0x2d, 0x50, 0xa9, 0xaa,
	0xda, 0xca, 0x12, 0xa9, 0xd4, 0xfb, 0xc2, 0x45, 0xd5, 0x96, 0x2b, 0xa8, 0x7a, 0xbb, 0x5a, 0xdb,
	0x13, 0xb3, 0xd4, 0x3f, 0xb0, 0xbb, 0x4e, 0xc8, 0x5b, 0xf5, 0x81, 0x7a, 0xd7, 0x47, 0xe8, 0x0b,
	0x74, 0xbc, 0x76, 0x7e, 0x4b, 0xae, 0x6c, 0xcf, 0x39, 0x33, 0x9e, 0x39, 0x73, 0x76, 0xc1, 0x0a,
	0x1f, 0xfc, 0x7b, 0x99, 0xeb, 0xfc, 0xec, 0xf7, 0x06, 0xb4, 0x2f, 0xf3, 0x6c, 0x28, 0x62, 0xa7,
	0x07, 0x9d, 0x11, 0x4a, 0x25, 0xf2, 0xcc, 0x6d, 0x9c, 0x34, 0x5e, 0xb7, 0xca, 0x40, 0xf8, 0xc0,
	0x32, 0x9e, 0xa2, 0xbb, 0x41, 0x81, 0xae, 0xf3, 0x0c, 0xba, 0xc4, 0x10, 0x43, 0x41, 0x34, 0xb7,
	0x49, 0x21, 0x7b, 0x00, 0xfe, 0xf7, 0x69, 0xc4, 0xd9, 0x87, 0x6d, 0xe2, 0x2b, 0xcd, 0x75, 0xa1,
	0x58, 0x21, 0x13, 0x77, 0xd3, 0x64, 0x9d, 0xc2, 0xd1, 0xad, 0x88, 0x90, 0x49, 0x1c, 0x32, 0x91,
	0xb1, 0x30, 0x4f, 0x53, 0xa1, 0x35, 0x46, 0x2c, 0x55, 0xb1, 0xdb, 0x22, 0x8a, 0xe5, 0x78, 0xe0,
	0x54, 0x61, 0x16, 0x14, 0x52, 0x69, 0x16, 0x61, 0xc2, 0x27, 0x6e, 0xdb, 0x74, 0xe1, 0xc2, 0x7f,
	0x29, 0x7f, 0x64, 0x8b, 0xb8, 0xdb, 0x31, 0x08, 0xfd, 0x8f, 0xea, 0xd1, 0x1c, 0x51, 0x11, 0xea,
	0xb2, 0x6d, 0xcb, 0x14, 0x3b, 0x06, 0x4b, 0x0a, 0xd4, 0x23, 0x4c, 0x22, 0xb7, 0x6b, 0x9a, 0xec,
	0xfa, 0xd7, 0x75, 0xc0, 0x39, 0x84, 0x76, 0x8c, 0x52, 0x0a, 0xed, 0xf6, 0x0c, 0xd4, 0xf1, 0x3f,
	0x99, 0x4f, 0xa7, 0x0f, 0x5b, 0x31, 0xd5, 0x97, 0x78, 0x9f, 0x9b, 0xde, 0xc1, 0xf4, 0xee, 0x00,
	0x68, 0x2e, 0x63, 0x2c, 0x81, 0xa1, 0x6b, 0x9b, 0x18, 0x31, 0xd5, 0x28, 0x9b, 0x33, 0xb7, 0x4c,
	0xf4, 0x18, 0xf6, 0x22, 0xc9, 0x45, 0x26, 0xb2, 0xb8, 0x94, 0x40, 0x6a, 0xa6, 0x05, 0x09, 0xb7,
	0x5d, 0x82, 0x67, 0x3e, 0x58, 0xb3, 0x0e, 0x6c, 0x68, 0x96, 0x59, 0x0d, 0x93, 0x45, 0x23, 0x50,
	0xff, 0x77, 0x18, 0xd2, 0x64, 0x5c, 0xa1, 0x22, 0xa1, 0x9b, 0xc4, 0x7f, 0x09, 0xed, 0xba, 0xad,
	0x23, 0xd8, 0x25, 0x4d, 0x6b, 0xd5, 0x23, 0x96, 0xf0, 0x00, 0xeb, 0xdc, 0xb3, 0x3f, 0x6d, 0xe8,
	0xce, 0xc5, 0xff, 0x00, 0xdb, 0x12, 0x47, 0x02, 0xc7, 0x28, 0x59, 0x12, 0xeb, 0xd4, 0x90, 0xec,
	0xc1, 0x8b, 0xf9, 0x7e, 0xfc, 0xeb, 0x1a, 0xbf, 0x22, 0x78, 0x1a, 0xa5, 0x3c, 0x5b, 0x4b, 0xc4,
	0x7a, 0x6d, 0x66, 0xd1, 0xf6, 0xe0, 0x74, 0x21, 0xeb, 0x1b, 0xa1, 0x37, 0x06, 0x5c, 0xca, 0x7b,
	0x03, 0x1d, 0x2d, 0x27, 0xec, 0x2e, 0x0f, 0x6a, 0x27, 0x1c, 0x2d, 0xe5, 0x4c, 0xbe, 0xe4, 0xc1,
	0x8c, 0xfb, 0x0e, 0x2c, 0x25, 0x62, 0x5a, 0x7d, 0xc2, 0x8d, 0x27, 0xec, 0x81, 0xb7, 0x40, 0xbe,
	0x21, 0xe8, 0xf2, 0xea, 0xe3, 0x34, 0xe0, 0x8d, 0xa1, 0xff, 0x64, 0xa7, 0x07, 0xb0, 0x33, 0xf5,
	0x0e, 0x8d, 0x28, 0xc8, 0x06, 0x33, 0x0d, 0x4b, 0x83, 0x8c, 0x39, 0xad, 0x4f, 0x61, 0x58, 0xcd,
	0xd0, 0x72, 0xf6, 0xc0, 0xce, 0x72, 0x23, 0x85, 0x31, 0x5a, 0x73, 0xbe, 0xa5, 0x09, 0x93, 0x45,
	0xc6, 0x78, 0x18, 0xa2, 0x52, 0x55, 0x21, 0x63, 0x54, 0xef, 0x1c, 0x0e, 0xd6, 0x0c, 0x7b, 0x08,
	0xbd, 0x05, 0x91, 0xd8, 0x6c, 0x7f, 0xde, 0xaf, 0x26, 0xec, 0xac, 0x0c, 0x7b, 0x0e, 0x9d, 0xa0,
	0x08, 0x7f, 0xa0, 0x56, 0xc4, 0x69, 0xd2, 0xac, 0x27, 0x6b, 0x85, 0xf1, 0x2f, 0x0c, 0xd1, 0xf9,
	0x0c, 0xfd, 0x5a, 0x4b, 0x72, 0x55, 0xf9, 0x16, 0x9a, 0x13, 0x59, 0x2f, 0xe3, 0xed, 0xfa, 0xfc,
	0xea, 0xf3, 0xba, 0xcc, 0xa9, 0x0e, 0xb1, 0x77, 0x05, 0x9d, 0x8b, 0x42, 0x24, 0x11, 0x35, 0xb2,
	0x05, 0x9b, 0xe6, 0xec, 0x36, 0xa6, 0xae, 0xd5, 0x52, 0xc4, 0x64, 0x7e, 0xb2, 0x51, 0x30, 0x99,
	0x9d, 0xe8, 0x7d, 0x7c, 0xbc, 0xa7, 0x72, 0x29, 0x66, 0x9a, 0xd1, 0x4b, 0x48, 0x4f, 0x1e, 0xa3,
	0x51, 0x64, 0xc3, 0xfb, 0x0a, 0xed, 0xba, 0xc5, 0xe5, 0x62, 0xef, 0xc1, 0x0a, 0xaa, 0xbf, 0x54,
	0x8e, 0x5d, 0x75, 0xcc, 0xca, 0x90, 0x86, 0xe9, 0xfd, 0x6c, 0xc0, 0xee, 0x3f, 0x0d, 0x97, 0x1b,
	0x59, 0x9e, 0xfd, 0xa1, 0xc8, 0x35, 0xaf, 0x6f, 0x20, 0xba, 0x17, 0xe2, 0x24, 0x0f, 0x78, 0xb2,
	0x84, 0x55, 0xfb, 0xfd, 0x1f, 0xfa, 0x43, 0x2e, 0x92, 0x42, 0x62, 0x0d, 0x8e, 0x51, 0xc4, 0xb7,
	0xda, 0x2c, 0xba, 0xe5, 0xbc, 0x82, 0xe7, 0x5a, 0xf2, 0x4c, 0x89, 0x72, 0xae, 0x27, 0x79, 0x9b,
	0xd3, 0x2a, 0xe5, 0x39, 0xcd, 0x0b, 0xbd, 0x8c, 0x96, 0xf7, 0x52, 0xcb, 0xdb, 0x85, 0xde, 0x8a,
	0x3b, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x92, 0xf5, 0x6a, 0x3b, 0x05, 0x00, 0x00,
}
