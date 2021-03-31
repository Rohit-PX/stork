// +build integrationtest

package integrationtest

import (
	"testing"

	"github.com/portworx/sched-ops/k8s/core"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	"github.com/portworx/torpedo/drivers/scheduler"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func restartStorkDestTest(
	t *testing.T,
) {
	// Create myqsl app deployment
	appCtx := createApp(t, appKey+"-restart-stork-dest")

	backupLocationName := appKey + "-backup-location"
	// Create backup location
	backupLocation, err := createBackupLocation(t, backupLocationName, appCtx.GetID(), defaultBackupLocation, defaultSecretName)
	require.NoError(t, err, "Error creating backuplocation")

	logrus.Infof("All Apps created %v. Starting backup.", appCtx.GetID())

	// Create first application backup
	appBackup, err := createApplicationBackupWithAnnotation(t, appKey+"-backup", appCtx.GetID(), backupLocation)

	// Wait for backup to start
	err = waitForAppBackupToStart(appBackup.Name, appBackup.Namespace, applicationBackupScheduleRetryTimeout)
	require.NoError(t, err, "Application backup %s failed to start", appBackup)

	err = setDestinationKubeConfig()
	require.NoError(t, err, "failed to set kubeconfig to destination cluster: %v", err)

	err = deleteStorkPods()
	require.NoError(t, err, "failed to delete  stork pods")

	err = waitForAppBackupCompletion(appBackup.Name, appBackup.Namespace, applicationBackupScheduleRetryTimeout)
	require.NoError(t, err, "Application backup %s in namespace %s failed.", appBackup.Name, appBackup.Namespace)

	logrus.Infof("Application backup successfully completed.")

	// Get application backup object to be used for validation of cloud delete later
	appBackup, err = storkops.Instance().GetApplicationBackup(appBackup.Name, appBackup.Namespace)
	require.NoError(t, err, "Failed to get application backup: %s in namespace: %s, for cloud deletion validation", appBackup.Name, appBackup.Namespace, err)

	firstBackupPath := appBackup.Status.BackupPath

	err = deleteAndWaitForBackupDeletion(appBackup.Namespace)
	require.NoError(t, err, "Backups not deleted: %v", err)

	err = storkops.Instance().DeleteBackupLocation(backupLocation.Name, backupLocation.Namespace)
	require.NoError(t, err, "Failed to delete backuplocation: %s", backupLocation.Name, err)

	validateBackupDeletionFromObjectstore(t, backupLocation, firstBackupPath)
	destroyAndWait(t, []*scheduler.Context{appCtx})
}

func deleteStorkPods() error {
	// Get Stork pods on destination cluster and restart them
	storkPods, err := core.Instance().GetPods("kube-system", map[string]string{"name": storkPodLabel})
	if err != nil {
		return err
	}

	err = core.Instance().DeletePods(storkPods.Items, false)
	if err != nil {
		return err
	}
	return nil

}
