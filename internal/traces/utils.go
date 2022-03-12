package traces

import (
	"os"

	"github.com/filecoin-project/bacalhau/internal/logger"
	"github.com/filecoin-project/bacalhau/internal/system"
	"github.com/filecoin-project/bacalhau/internal/types"
)

// return 2 lists of job ids - correct and incorrect
func ProcessResults(job *types.Job, data *[]system.ResultsList) ([]string, []string, error) {
	// TODO: load the job so we can get at Deal.Tolerance
	clustered := TraceCollection{
		Traces:    []Trace{},
		Tolerance: job.Deal.Tolerance,
	}

	for _, row := range *data {
		resultsFolder, err := system.GetSystemDirectory(row.Folder)
		if err != nil {
			return []string{}, []string{}, err
		}

		if _, err := os.Stat(resultsFolder); os.IsNotExist(err) {
			logger.Info("Results folder does not exist, continuing.")
			continue
		}
		clustered.Traces = append(clustered.Traces, Trace{
			ResultId: row.Cid,
			Filename: resultsFolder + "/metrics.log",
		})
	}

	// these are lists of CIDs of the results
	correctGroup, incorrectGroup, err := clustered.Cluster()

	if err != nil {
		return []string{}, []string{}, err
	}

	return correctGroup, incorrectGroup, err
}
