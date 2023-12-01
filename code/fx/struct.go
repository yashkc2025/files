package fx

type GreenHouse struct {
	Jobs []struct {
		AbsoluteURL    string `json:"absolute_url"`
		DataCompliance []struct {
			Type                      string      `json:"type"`
			RequiresConsent           bool        `json:"requires_consent"`
			RequiresProcessingConsent bool        `json:"requires_processing_consent"`
			RequiresRetentionConsent  bool        `json:"requires_retention_consent"`
			RetentionPeriod           interface{} `json:"retention_period"`
		} `json:"data_compliance"`
		InternalJobID int `json:"internal_job_id"`
		Location      struct {
			Name string `json:"name"`
		} `json:"location"`
		Metadata      interface{} `json:"metadata"`
		ID            int         `json:"id"`
		UpdatedAt     string      `json:"updated_at"`
		RequisitionID string      `json:"requisition_id"`
		Title         string      `json:"title"`
		Content       string      `json:"content"`
		Departments   []struct {
			ID       int           `json:"id"`
			Name     string        `json:"name"`
			ChildIds []interface{} `json:"child_ids"`
			ParentID interface{}   `json:"parent_id"`
		} `json:"departments"`
		Offices []struct {
			ID       int           `json:"id"`
			Name     string        `json:"name"`
			Location string        `json:"location"`
			ChildIds []interface{} `json:"child_ids"`
			ParentID int           `json:"parent_id"`
		} `json:"offices"`
	} `json:"jobs"`
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
}

type Workable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Jobs        []struct {
		Title          string `json:"title"`
		Shortcode      string `json:"shortcode"`
		Code           string `json:"code"`
		EmploymentType string `json:"employment_type"`
		Telecommuting  bool   `json:"telecommuting"`
		Department     string `json:"department"`
		URL            string `json:"url"`
		Shortlink      string `json:"shortlink"`
		ApplicationURL string `json:"application_url"`
		PublishedOn    string `json:"published_on"`
		CreatedAt      string `json:"created_at"`
		Country        string `json:"country"`
		City           string `json:"city"`
		State          string `json:"state"`
		Education      string `json:"education"`
		Experience     string `json:"experience"`
		Function       string `json:"function"`
		Industry       string `json:"industry"`
		Locations      []struct {
			Country     string `json:"country"`
			CountryCode string `json:"countryCode"`
			City        string `json:"city"`
			Region      string `json:"region"`
			Hidden      bool   `json:"hidden"`
		} `json:"locations"`
		Description string `json:"description"`
	} `json:"jobs"`
}
