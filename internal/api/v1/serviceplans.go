package v1

import (
	"encoding/json"
	"net/http"

	"github.com/epinio/epinio/helpers/kubernetes"
	"github.com/epinio/epinio/internal/services"
	"github.com/julienschmidt/httprouter"
)

type ServicePlansController struct {
}

func (spc ServicePlansController) Index(w http.ResponseWriter, r *http.Request) APIErrors {
	params := httprouter.ParamsFromContext(r.Context())
	serviceClassName := params.ByName("serviceclass")

	cluster, err := kubernetes.GetCluster()
	if err != nil {
		return InternalError(err)
	}

	serviceClass, err := services.ClassLookup(cluster, serviceClassName)
	if err != nil {
		return InternalError(err)
	}

	if serviceClass == nil {
		return ServiceClassIsNotKnown(serviceClassName)
	}
	servicePlans, err := serviceClass.ListPlans()
	if err != nil {
		return InternalError(err)
	}

	js, err := json.Marshal(servicePlans)
	if err != nil {
		return InternalError(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		return InternalError(err)
	}

	return nil
}
