package validation

// swagger:parameters CreateTenantRequest
type CreateTenantRequest struct {
	// required: true
	Task string `form:"task" json:"task" xml:"task"  binding:"required,min=1,max=300"`
	// required: true
	Status string `form:"status" json:"status" xml:"status"  binding:"required,oneof=active inactive"`
}

// swagger:parameters UpdateTenantRequest
type UpdateTenantRequest struct {
	// required: true
    Task string `form:"task" json:"task" xml:"task"  binding:"required,min=1,max=300"`
    // required: true
    Status string `form:"status" json:"status" xml:"status"  binding:"required,oneof=active inactive"`
}
