package external_miner

import (
	"fmt"

	"github.com/kurtosis-tech/eth2-merge-kurtosis-module/kurtosis-module/impl/participant_network/el"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/enclaves"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/services"
	"github.com/kurtosis-tech/stacktrace"
)

const (
	imageName = "marioevz/external_miner:latest"
	serviceId = "external-miner"
)

// TODO upgrade the spammer to be able to take in multiple EL addreses
func LaunchExternalMiner(enclaveCtx *enclaves.EnclaveContext, elClientCtxs []*el.ELClientContext, ttd uint64) error {
	containerConfigSupplier := getContainerConfigSupplier(elClientCtxs, ttd)

	_, err := enclaveCtx.AddService(serviceId, containerConfigSupplier)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred adding the external miner service")
	}

	return nil
}

func getContainerConfigSupplier(elClientCtxs []*el.ELClientContext, ttd uint64) func(string, *services.SharedPath) (*services.ContainerConfig, error) {

	return func(privateIpAddr string, sharedDir *services.SharedPath) (*services.ContainerConfig, error) {
		cmd := []string{
			"--delay",
			"TTD,600",
			"--ttd",
			fmt.Sprintf("%d", ttd),
		}
		for _, elClientCtx := range elClientCtxs {
			cmd = append(cmd, "--rpc")
			cmd = append(cmd, fmt.Sprintf("http://%v:%v", elClientCtx.GetIPAddress(), elClientCtx.GetRPCPortNum()))
		}
		result := services.NewContainerConfigBuilder(
			imageName,
		).WithCmdOverride(cmd).Build()
		return result, nil
	}
}
