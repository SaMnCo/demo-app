package model

type Policy struct {
        Id string `json:"policyNumber"`
        Name string  `json:"name"`
        ServiceVersion string `json:"serviceVersion"`
        RiskAddress string `json:"riskAddress"`
}

func (policy *Policy) ToString() string {
    return policy.Id + " \n" + policy.Name + " \n" + policy.ServiceVersion + " \n" + policy.RiskAddress
}