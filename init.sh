# 1. Crear la estructura base del proyecto
# mkdir -p blue-whale-platform
# cd blue-whale-platform

# 2. Crear directorios principales
mkdir -p .github/{workflows,} \
        backend/services/{idp,tenant-service,api-gateway}/{cmd,internal,config,migrations,test,features} \
        backend/shared/{common,pkg} \
        deployments/{docker,kubernetes,terraform}/{idp,tenant-service,api-gateway} \
        docs/{architecture/{diagrams,decisions},api/openapi,development} \
        scripts tools web/admin-panel mobile/{android,ios}

# IDP Service
cd backend/services/idp
go mod init github.com/mikemonzo/blue-whale-platform/backend/services/idp
mkdir -p internal/{application/{commands,queries,services},domain/{model,repository,service},infrastructure/{persistence,auth,api}}
# touch cmd/main.go config/{config.go,config.yaml}

# Crear estructura BDD para IDP
mkdir -p features/{authentication,authorization,tenant_management,user_management}/{steps,support}
# touch features/authentication/{login,logout,token_validation}.feature
# touch features/authorization/{rbac,permissions}.feature
# touch features/tenant_management/{create_tenant,assign_user}.feature
# touch features/user_management/{create_user,update_user}.feature

# Tenant Service
cd ../tenant-service
go mod init github.com/mikemonzo/blue-whale-platform/backend/services/tenant-service
mkdir -p internal/{application/{commands,queries,services},domain/{model,repository,service},infrastructure/{persistence,auth,api}}
#  touch cmd/main.go config/{config.go,config.yaml}

# Crear estructura BDD para Tenant Service
mkdir -p features/{tenant_management,tenant_configuration,tenant_metrics}/{steps,support}
# touch features/tenant_management/{create_tenant,update_tenant,delete_tenant}.feature
# touch features/tenant_configuration/{configure_tenant,tenant_settings}.feature
# touch features/tenant_metrics/{usage_metrics,billing_metrics}.feature

# API Gateway
cd ../api-gateway
go mod init github.com/mikemonzo/blue-whale-platform/backend/services/api-gateway
mkdir -p internal/{application/{commands,queries,services},domain/{model,repository,service},infrastructure/{persistence,auth,api}}
# touch cmd/main.go config/{config.go,config.yaml}

# Crear estructura BDD para API Gateway
mkdir -p features/{routing,authentication,authorization}/{steps,support}
# touch features/routing/{route_management,load_balancing}.feature
# touch features/authentication/{token_validation,session_management}.feature
# touch features/authorization/{access_control,rate_limiting}.feature

# 5. Añadir godog a cada servicio
cd ../idp
go get github.com/cucumber/godog@v0.12.6
cd ../tenant-service
go get github.com/cucumber/godog@v0.12.6
cd ../api-gateway
go get github.com/cucumber/godog@v0.12.6

# 6. Volver a la raíz y crear el workspace
pwd
cd ../../
go work init
go work use ./shared/common ./shared/pkg ./services/idp ./services/tenant-service ./services/api-gateway