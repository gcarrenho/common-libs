// this package contains the services. it are our entry points to the core and each one of them implements the corresponding port
package services

import (
	"encoding/json"
	"io"

	"github.azc.ext.hp.com/LFP-Tropos/udc-common-libs/src/internal/models"

	"github.com/rs/zerolog"
)

type LoggingService struct {
	log *zerolog.Logger
}

func NewLoggingService(w io.Writer) *LoggingService {
	zerolog := zerolog.New(w)
	return &LoggingService{
		log: &zerolog,
	}
}

func (logsrv *LoggingService) Panic(logging *models.Logging) {
	logMap := structToMap(logging)
	log := logsrv.log.With().Fields(logMap).Logger()
	log.Panic().Msg(logging.Message)
}

func (logsrv *LoggingService) Fatal(logging *models.Logging) {
	logMap := structToMap(logging)
	log := logsrv.log.With().Fields(logMap).Logger()
	log.Fatal().Msg(logging.Message)
}

func (logsrv *LoggingService) Error(logging *models.Logging) {
	logMap := structToMap(logging)
	log := logsrv.log.With().Fields(logMap).Logger()
	log.Error().Msg(logging.Message)
}

func (logsrv *LoggingService) Warn(logging *models.Logging) {
	logMap := structToMap(logging)
	log := logsrv.log.With().Fields(logMap).Logger()
	log.Warn().Msg(logging.Message)

}

func (logsrv *LoggingService) Info(logging *models.Logging) {
	logMap := structToMap(logging)
	log := logsrv.log.With().Fields(logMap).Logger()
	log.Info().Msg(logging.Message)
}

func (logsrv *LoggingService) Debug(logging *models.Logging) {
	logMap := structToMap(logging)
	log := logsrv.log.With().Fields(logMap).Logger()
	log.Debug().Msg(logging.Message)
}

func structToMap(logging *models.Logging) map[string]interface{} {
	logMap := make(map[string]interface{})
	j, _ := json.Marshal(logging)
	json.Unmarshal(j, &logMap)
	return logMap
}
