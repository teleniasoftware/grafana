/*Package api contains base API implementation of unified alerting
 *
 *Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 *Do not manually edit these files, please find ngalert/api/swagger-codegen/ for commands on how to generate them.
 */
package api

import (
	"github.com/go-macaron/binding"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/models"
	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
)

type RulerApiService interface {
	RouteDeleteNamespaceRulesConfig(*models.ReqContext) response.Response
	RouteDeleteRuleGroupConfig(*models.ReqContext) response.Response
	RouteGetNamespaceRulesConfig(*models.ReqContext) response.Response
	RouteGetRulegGroupConfig(*models.ReqContext) response.Response
	RouteGetRulesConfig(*models.ReqContext) response.Response
	RoutePostNameRulesConfig(*models.ReqContext, apimodels.PostableRuleGroupConfig) response.Response
}

func (api *API) RegisterRulerApiEndpoints(srv RulerApiService) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Delete(toMacaronPath("/api/ruler/{Recipient}/api/v1/rules/{Namespace}"), routing.Wrap(srv.RouteDeleteNamespaceRulesConfig))
		group.Delete(toMacaronPath("/api/ruler/{Recipient}/api/v1/rules/{Namespace}/{Groupname}"), routing.Wrap(srv.RouteDeleteRuleGroupConfig))
		group.Get(toMacaronPath("/api/ruler/{Recipient}/api/v1/rules/{Namespace}"), routing.Wrap(srv.RouteGetNamespaceRulesConfig))
		group.Get(toMacaronPath("/api/ruler/{Recipient}/api/v1/rules/{Namespace}/{Groupname}"), routing.Wrap(srv.RouteGetRulegGroupConfig))
		group.Get(toMacaronPath("/api/ruler/{Recipient}/api/v1/rules"), routing.Wrap(srv.RouteGetRulesConfig))
		group.Post(toMacaronPath("/api/ruler/{Recipient}/api/v1/rules/{Namespace}"), binding.Bind(apimodels.PostableRuleGroupConfig{}), routing.Wrap(srv.RoutePostNameRulesConfig))
	}, middleware.ReqSignedIn)
}
