package handler

import (
    "fmt"
)

func errParamRequired(name, typ string) error {
    return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateWorkRequest struct {
    Role     string `json:"role"`
    Company  string `json:"company"`
    Location string `json:"location"`
    Remote   *bool  `json:"remote"`
    Link     string `json:"link"`
    Salary   int64  `json:"salary"`
}

func (r *CreateWorkRequest) Validate() error {
    if r.Role == "" {
        return errParamRequired("role", "string")
    }
    if r.Company == "" {
        return errParamRequired("company", "string")
    }
    if r.Location == "" {
        return errParamRequired("location", "string")
    }
    if r.Remote == nil {
        return errParamRequired("remote", "bool")
    }
    if r.Link == "" {
        return errParamRequired("link", "string")
    }
    if r.Salary <= 0 {
        return errParamRequired("salary", "int")
    }
    return nil
}

type UpdateWorkRequest struct {
    Role     string `json:"role"`
    Company  string `json:"company"`
    Location string `json:"location"`
    Remote   *bool  `json:"remote"`
    Link     string `json:"link"`
    Salary   int64  `json:"salary"`
}

func (r *UpdateWorkRequest) Validate() error {
    if r.Role != "" || r.Company != "" || r.Location != "" ||
        r.Remote != nil || r.Link != "" || r.Salary > 0 {
        return nil
    }
    return fmt.Errorf("at least one valid field must be provided")
}
