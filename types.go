package main

// Application ...
type Application struct {
	AppName   string            `json:"application,omitempty"`
	Globals   map[string]string `json:"globals,omitempty"`
	Pipelines []Pipeline        `json:"pipelines,omitempty"`
}

// Pipeline - main pipelineName struct
type Pipeline struct {
	Application          string        `json:"application,omitempty"`
	Name                 string        `json:"name,omitempty"`
	KeepWaitingPipelines bool          `json:"keepWaitingPipelines,omitempty"`
	LimitConcurrent      bool          `json:"limitConcurrent,omitempty"`
	Parallel             bool          `json:"parallel,omitempty"`
	ParameterConfigs     []interface{} `json:"parameterConfig,omitempty"`
	Triggers             []interface{} `json:"triggers,omitempty"`
	Stages               []interface{} `json:"stages,omitempty"`
}

// ParameterConfig
type ParameterConfig struct {
	Name        string      `json:"name,omitempty"`
	Default     string      `json:"default,omitempty"`
	Description string      `json:"description,omitempty"`
	HasOptions  interface{} `json:"hasOptions,omitempty"`
	Label       string      `json:"label,omitempty"`
	Options     []string    `json:"options,omitempty"`
	Pinned      interface{} `json:"pinned,omitempty"`
	Required    interface{} `json:"Required,omitempty"`
}

// Artifacts
// ArtifactGithub
type ArtifactGithub struct {
	ArtifactAccount string `json:"artifactAccount,omitempty"`
	Name            string `json:"name,omitempty"`
	Reference       string `json:"reference,omitempty"`
	Type            string `json:"type,omitempty"`
	Version         string `json:"version,omitempty"`
}

// Stages
// DeployManifest ...
type DeployManifest struct {
	Name                     string            `json:"name,omitempty"`
	Type                     string            `json:"type,omitempty"`
	RefId                    string            `json:"refId,omitempty"`
	RequisiteStageRefIds     []string          `json:"requisiteStageRefIds,omitempty"`
	SkipExpressionEvaluation bool              `json:"skipExpressionEvaluation,omitempty"`
	Account                  string            `json:"account,omitempty"`
	CloudProvider            string            `json:"cloudProvider,omitempty"`
	ManifestArtifact         interface{}       `json:"manifestArtifact,omitempty"`
	Manifests                interface{}       `json:"manifests,omitempty"`
	Moniker                  map[string]string `json:"moniker,omitempty"`
	TrafficMgmt              `json:"trafficManagement,omitempty"`
}

// DeleteManifest ...
type DeleteManifest struct {
	Name                     string                 `json:"name,omitempty"`
	Type                     string                 `json:"type,omitempty"`
	RefId                    string                 `json:"refId,omitempty"`
	RequisiteStageRefIds     []string               `json:"requisiteStageRefIds,omitempty"`
	SkipExpressionEvaluation bool                   `json:"skipExpressionEvaluation,omitempty"`
	Account                  string                 `json:"account,omitempty"`
	CloudProvider            string                 `json:"cloudProvider,omitempty"`
	ManifestName             string                 `json:"manifestName,omitempty"`
	Mode                     string                 `json:"mode,omitempty"`
	Options                  map[string]interface{} `json:"options,omitempty"`
	Manifests                interface{}
	Location                 string `json:"location,omitempty"`
	StageEnabled             `json:"stageEnabled,omitempty"`
	TrafficMgmt              `json:"trafficManagement,omitempty"`
}

//ManualJudgement
type ManualJudgement struct {
	Name                 string            `json:"name,omitempty"`
	Type                 string            `json:"type,omitempty"`
	RefId                string            `json:"refId,omitempty"`
	RequisiteStageRefIds []string          `json:"requisiteStageRefIds,omitempty"`
	FailPipeline         bool              `json:"failPipeline,omitempty"`
	Instructions         string            `json:"instructions,omitempty"`
	JudgmentInputs       map[string]string `json:"judgementInputs,omitempty"`
	Notifications        []Notification    `json:"notifications,omitempty"`
	SendNotifications    bool              `json:"sendNotifications,omitempty"`
}

//Terraform
// Going to try generic TF first to see if all TF stages will work here
type Terraform struct {
	Name                  string            `json:"name,omitempty"`
	Type                  string            `json:"type,omitempty"`
	RefId                 string            `json:"refId,omitempty"`
	RequisiteStageRefIds  []string          `json:"requisiteStageRefIds,omitempty"`
	Action                string            `json:"action,omitempty"`
	Artifacts             []interface{}     `json:"artifacts,omitempty"`
	CompleteOtherBranches bool              `json:"completeOtherBranches,omitempty"`
	ContinuePipeline      bool              `json:"continuePipeline,omitempty"`
	ExpectedArtifacts     []interface{}     `json:"expectedArtifacts,omitempty"`
	FailPipeline          bool              `json:"failPipeline,omitempty"`
	Overrides             map[string]string `json:"Overrides,omitempty"`
	Profile               string            `json:"profile,omitempty"`
	Targets               []string          `json:"targets,omitempty"`
	TerraformVersion      string            `json:"terraformVersion,omitempty"`
	Workspace             string            `json:"workspace,omitempty"`
}

//TfApply
type TfApply struct {
	Name                          string            `json:"name,omitempty"`
	Type                          string            `json:"type,omitempty"`
	RefId                         string            `json:"refId,omitempty"`
	RequisiteStageRefIds          []string          `json:"requisiteStageRefIds,omitempty"`
	Action                        string            `json:"action,omitempty"`
	Artifacts                     []interface{}     `json:"artifacts,omitempty"`
	CompleteOtherBranches         bool              `json:"completeOtherBranches,omitempty"`
	CompleteOtherBranchesThenFail bool              `json:"completeOtherBranches,omitempty"`
	ContinuePipeline              bool              `json:"continuePipeline,omitempty"`
	ExpectedArtifacts             []interface{}     `json:"expectedArtifacts,omitempty"`
	FailPipeline                  bool              `json:"failPipeline,omitempty"`
	Overrides                     map[string]string `json:"Overrides,omitempty"`
	Profile                       string            `json:"profile,omitempty"`
	Targets                       []string          `json:"targets,omitempty"`
	TerraformVersion              string            `json:"terraformVersion,omitempty"`
	Workspace                     string            `json:"workspace,omitempty"`
}

//TfDestroy
type TfDestroy struct {
	Name                          string            `json:"name,omitempty"`
	Type                          string            `json:"type,omitempty"`
	RefId                         string            `json:"refId,omitempty"`
	RequisiteStageRefIds          []string          `json:"requisiteStageRefIds,omitempty"`
	Action                        string            `json:"action,omitempty"`
	Artifacts                     []interface{}     `json:"artifacts,omitempty"`
	CompleteOtherBranches         bool              `json:"completeOtherBranches,omitempty"`
	CompleteOtherBranchesThenFail bool              `json:"completeOtherBranches,omitempty"`
	ContinuePipeline              bool              `json:"continuePipeline,omitempty"`
	ExpectedArtifacts             []interface{}     `json:"expectedArtifacts,omitempty"`
	FailPipeline                  bool              `json:"failPipeline,omitempty"`
	Overrides                     map[string]string `json:"Overrides,omitempty"`
	Profile                       string            `json:"profile,omitempty"`
	Targets                       []string          `json:"targets,omitempty"`
	TerraformVersion              string            `json:"terraformVersion,omitempty"`
	Workspace                     string            `json:"workspace,omitempty"`
}

type Notification struct {
	Address string                 `json:"address,omitempty"`
	Level   string                 `json:"level,omitempty"`
	Message map[string]interface{} `json:"message,omitempty"`
	Type    string                 `json:"type,omitempty"`
	When    []string               `json:"when,omitempty"`
}

type StageEnabled struct {
	Expression string `json:"expression,omitempty"`
	Type       string `json:"type,omitempty"`
}

type TrafficMgmt struct {
	Enabled bool               `json:"enabled,omitempty"`
	Options TrafficMgmtOptions `json:"options,omitempty"`
}

type TrafficMgmtOptions struct {
	EnableTraffic bool     `json:"enableTraffic,omitempty"`
	Services      []string `json:"services,omitempty"`
}

// Triggers
type TriggerGeneric struct {
	Type string `json:"type,omitempty"`
}

// TriggerDocker
type TriggerDocker struct {
	Account             string   `json:"account,omitempty"`
	Enabled             interface{}    `json:"enabled,omitempty"`
	ExpectedArtifactIds []string `json:"expectedArtifactIds,omitempty"`
	Organization        string   `json:"organization,omitempty"`
	Registry            string   `json:"registry,omitempty"`
	Repository          string   `json:"repository,omitempty"`
	Tag                 string   `json:"tag,omitempty"`
	Type                string   `json:"type,omitempty"`
}

//TriggerGitHub
type TriggerGithub struct {
	Branch              string   `json:"branch,omitempty"`
	Enabled             bool     `json:"enabled,omitempty"`
	ExpectedArtifactIds []string `json:"expectedArtifactIds,omitempty"`
	Project             string   `json:"project,omitempty"`
	Secret              string   `json:"secret,omitempty"`
	Slug                string   `json:"slug,omitempty"`
	Source              string   `json:"source,omitempty"`
	Type                string   `json:"type,omitempty"`
}
