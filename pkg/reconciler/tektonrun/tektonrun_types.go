package tektonrun

const (
	ParamSourceURL      = "shp-source-url"
	ParamSourceRevision = "shp-source-revision"
	ParamOutputImage    = "shp-output-image"
)

// ExtraFields carry on metainformation to link a given Tekton Run object with Shipwright.
type ExtraFields struct {
	BuildRunName string `json:"buildRunName,omitempty"`
}

// IsEmpty checks if the BuildRunName is defined.
func (s *ExtraFields) IsEmpty() bool {
	return s.BuildRunName == ""
}
