package serviceclient

import (
	"context"
	"givc/api/admin"
	systemd_api "givc/api/systemd"
	givc_grpc "givc/internal/pkgs/grpc"
	"givc/internal/pkgs/types"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func GetRemoteStatus(cfg *types.EndpointConfig, unitName string) (*types.UnitStatus, error) {

	// Setup and dial GRPC client
	var conn *grpc.ClientConn
	conn, err := givc_grpc.NewDial(cfg)
	if err != nil {
		log.Errorf("Cannot create grpc client: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create client
	client := systemd_api.NewUnitControlServiceClient(conn)
	if client == nil {
		log.Errorf("Failed to create 'NewUnitControlServiceClient'")
		return nil, err
	}

	// Get unit status
	request := systemd_api.UnitRequest{
		UnitName: unitName,
	}
	ctx := context.Background()
	resp, err := client.GetUnitStatus(ctx, &request)
	if err != nil {
		log.Errorf("Not the response we hoped for: %v", err)
		return nil, err
	}

	response := &types.UnitStatus{
		Name:        resp.UnitStatus.Name,
		Description: resp.UnitStatus.Description,
		LoadState:   resp.UnitStatus.LoadState,
		ActiveState: resp.UnitStatus.ActiveState,
		SubState:    resp.UnitStatus.SubState,
		Path:        string(resp.UnitStatus.Path),
	}

	return response, nil
}

func RegisterRemoteService(cfg *types.EndpointConfig, reg *admin.RegistryRequest) (*admin.RegistryResponse, error) {

	// Setup and dial GRPC client
	var conn *grpc.ClientConn
	conn, err := givc_grpc.NewDial(cfg)
	if err != nil {
		log.Errorf("Cannot create grpc client: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create admin client
	client := admin.NewAdminServiceClient(conn)
	if client == nil {
		log.Errorf("Failed to create new admin client")
		return nil, err
	}

	// Send registry request
	log.Infof("Sending request: %v", reg)
	ctx := context.Background()
	resp, err := client.RegisterService(ctx, reg)
	if err != nil {
		log.Errorf("Not the response we hoped for: %v", err)
		return nil, err
	}
	log.Infoln(resp)

	return resp, nil
}

func StartRemoteService(cfg *types.EndpointConfig, unitName string) (*systemd_api.UnitResponse, error) {

	// Setup and dial GRPC client
	var conn *grpc.ClientConn
	conn, err := givc_grpc.NewDial(cfg)
	if err != nil {
		log.Errorf("Cannot create grpc client: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create client
	client := systemd_api.NewUnitControlServiceClient(conn)
	if client == nil {
		log.Errorf("Failed to create 'NewUnitControlServiceClient'")
		return nil, err
	}

	// Start unit
	request := systemd_api.UnitRequest{
		UnitName: unitName,
	}
	ctx := context.Background()
	resp, err := client.StartUnit(ctx, &request)
	if err != nil {
		log.Errorf("Not the response we hoped for: %v", err)
		return nil, err
	}

	log.Infoln(resp)
	return resp, nil
}

func StartRemoteApplication(cfg *types.EndpointConfig, unitName string) (*systemd_api.UnitResponse, error) {

	// Setup and dial GRPC client
	var conn *grpc.ClientConn
	conn, err := givc_grpc.NewDial(cfg)
	if err != nil {
		log.Errorf("Cannot create grpc client: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create client
	client := systemd_api.NewUnitControlServiceClient(conn)
	if client == nil {
		log.Errorf("Failed to create 'NewUnitControlServiceClient'")
		return nil, err
	}

	// Start unit
	request := systemd_api.UnitRequest{
		UnitName: unitName,
	}
	ctx := context.Background()
	resp, err := client.StartApplication(ctx, &request)
	if err != nil {
		log.Errorf("Not the response we hoped for: %v", err)
		return nil, err
	}

	log.Infoln(resp)
	return resp, nil
}

func PauseRemoteService(cfg *types.EndpointConfig, unitName string) (*systemd_api.UnitResponse, error) {

	// Setup and dial GRPC client
	var conn *grpc.ClientConn
	conn, err := givc_grpc.NewDial(cfg)
	if err != nil {
		log.Errorf("Cannot create grpc client: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create client
	client := systemd_api.NewUnitControlServiceClient(conn)
	if client == nil {
		log.Errorf("Failed to create 'NewUnitControlServiceClient'")
		return nil, err
	}

	// Start unit
	request := systemd_api.UnitRequest{
		UnitName: unitName,
	}
	ctx := context.Background()
	resp, err := client.FreezeUnit(ctx, &request)
	if err != nil {
		log.Errorf("Not the response we hoped for: %v", err)
		return nil, err
	}

	log.Infoln(resp)
	return resp, nil
}

func ResumeRemoteService(cfg *types.EndpointConfig, unitName string) (*systemd_api.UnitResponse, error) {

	// Setup and dial GRPC client
	var conn *grpc.ClientConn
	conn, err := givc_grpc.NewDial(cfg)
	if err != nil {
		log.Errorf("Cannot create grpc client: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create client
	client := systemd_api.NewUnitControlServiceClient(conn)
	if client == nil {
		log.Errorf("Failed to create 'NewUnitControlServiceClient'")
		return nil, err
	}

	// Start unit
	request := systemd_api.UnitRequest{
		UnitName: unitName,
	}
	ctx := context.Background()
	resp, err := client.UnfreezeUnit(ctx, &request)
	if err != nil {
		log.Errorf("Not the response we hoped for: %v", err)
		return nil, err
	}

	log.Infoln(resp)
	return resp, nil
}

func StopRemoteService(cfg *types.EndpointConfig, unitName string) (*systemd_api.UnitResponse, error) {

	// Setup and dial GRPC client
	var conn *grpc.ClientConn
	conn, err := givc_grpc.NewDial(cfg)
	if err != nil {
		log.Errorf("Cannot create grpc client: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create client
	client := systemd_api.NewUnitControlServiceClient(conn)
	if client == nil {
		log.Errorf("Failed to create 'NewUnitControlServiceClient'")
		return nil, err
	}

	// Start unit
	request := systemd_api.UnitRequest{
		UnitName: unitName,
	}
	ctx := context.Background()
	resp, err := client.StopUnit(ctx, &request)
	if err != nil {
		log.Errorf("Not the response we hoped for: %v", err)
		return nil, err
	}

	log.Infoln(resp)
	return resp, nil
}
