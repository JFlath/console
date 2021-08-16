package connect

import (
	"fmt"
	"github.com/cloudhut/common/rest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

// getMapValueOrString returns the map entry for the given key. If this entry does not exist it will return the
// passed fallback string.
func getMapValueOrString(m map[string]string, key string, fallback string) string {
	if val, exists := m[key]; exists {
		return val
	}

	return fallback
}

func (s *Service) getConnectClusterByName(clusterName string) (*ClientWithConfig, *rest.Error) {
	c, exists := s.ClientsByCluster[clusterName]
	if !exists {
		return nil, &rest.Error{
			Err:          fmt.Errorf("a client for the given cluster name does not exist"),
			Status:       http.StatusNotFound,
			Message:      "There's no configured cluster with the given connect cluster name",
			InternalLogs: []zapcore.Field{zap.String("cluster_name", clusterName)},
			IsSilent:     false,
		}
	}

	return c, nil
}
