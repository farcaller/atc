package migrations

import "github.com/BurntSushi/migration"

var Migrations = []migration.Migrator{
	InitialSchema,
	MoveSourceAndMetadataToVersionedResources,
	AddTypeToVersionedResources,
	RemoveTransitionalCurrentVersions,
	NonNullableVersionInfo,
	AddOneOffNameSequence,
	AddHijackURLToBuilds,
	AddTimestampsToBuilds,
	CreateLocks,
	AddBuildEvents,
	ReplaceBuildsAbortHijackURLsWithGuidAndEndpoint,
	ReplaceBuildEventsIDWithEventID,
	AddLocks,
	DropOldLocks,
	AddConfig,
	AddNameToBuildInputs,
	AddEngineAndEngineMetadataToBuilds,
	AddVersionToBuildEvents,
	AddCompletedToBuilds,
	AddWorkers,
	AddEnabledToBuilds,
	CreateEventIDSequencesForInFlightBuilds,
	AddResourceTypesToWorkers,
	AddPlatformAndTagsToWorkers,
	AddIdToConfig,
	ConvertJobBuildConfigToJobPlans,
	AddCheckErrorToResources,
	AddPausedToResources,
	AddPausedToJobs,
	CreateJobsSerialGroups,
	CreatePipes,
	RenameConfigToPipelines,
	RenamePipelineIDToVersionAddPrimaryKey,
	AddNameToPipelines,
	AddPipelineIDToResources,
	AddPipelineIDToJobs,
	AddPausedToPipelines,
	AddOrderingToPipelines,
	AddInputsDeterminedToBuilds,
	AddExplicitToBuildOutputs,
	AddLastCheckedToResources,
	AddLastTrackedToBuilds,
	AddLastScheduledToPipelines,
	AddCheckingToResources,
	AddUniqueConstraintToResources,
	RemoveSourceFromVersionedResources,
	AddIndexesToABunchOfStuff,
	DropLocks,
	AddBaggageclaimURLToWorkers,
	AddContainers,
	AddNameToWorkers,
	AddLastScheduledToBuilds,
	AddCheckTypeAndCheckSourceToContainers,
	AddStepLocationToContainers,
	AddVolumesAndCacheInvalidator,
	AddCompositeUniqueConstraintToVolumes,
	AddWorkingDirectoryToContainers,
	MakeContainerWorkingDirectoryNotNull,
	AddEnvVariablesToContainers,
	AddModifiedTimeToVersionedResourcesAndBuildOutputs,
	ReplaceStepLocationWithPlanID, //even though this is 56, it's meant to be here
	AddTeamsColumnToPipelinesAndTeamsTable,
	CascadePipelineDeletes,
	AddTeamIDToPipelineNameUniqueness,
	MakeVolumesExpiresAtNullable,
	AddAuthFieldsToTeams,
	AddAdminToTeams,
	MakeContainersLinkToPipelineIds,
	MakeContainersLinkToResourceIds,
	MakeContainersBuildIdsNullable,
	MakeContainersLinkToWorkerIds,
	RemoveVolumesWithExpiredWorkers,
	AddWorkerIDToVolumes,
	RemoveWorkerIds,
	AddAttemptsToContainers,
	AddStageToContainers,
	AddImageResourceVersions,
	MakeContainerIdentifiersUnique,
}
