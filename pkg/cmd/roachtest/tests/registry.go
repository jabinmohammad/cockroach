// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tests

import "github.com/cockroachdb/cockroach/pkg/cmd/roachtest/registry"

// RegisterTests registers all tests to the Registry. This powers `roachtest run`.
func RegisterTests(r registry.Registry) {
	registerAWSDMS(r)
	registerAcceptance(r)
	registerActiveRecord(r)
	registerAdmission(r)
	registerAllocator(r)
	registerAllocationBench(r)
	registerAlterPK(r)
	registerAsyncpg(r)
	registerAutoUpgrade(r)
	registerBackup(r)
	registerBackupMixedVersion(r)
	registerBackupNodeShutdown(r)
	registerBackupRestoreRoundTrip(r)
	registerCDC(r)
	registerCDCMixedVersions(r)
	registerCancel(r)
	registerChangeReplicasMixedVersion(r)
	registerClearRange(r)
	registerClockJumpTests(r)
	registerClockMonotonicTests(r)
	registerClusterToCluster(r)
	registerClusterReplicationResilience(r)
	registerClusterReplicationDisconnect(r)
	registerConnectionLatencyTest(r)
	registerCopy(r)
	registerCopyFrom(r)
	registerCostFuzz(r)
	registerDecommission(r)
	registerDecommissionBench(r)
	registerDiskFull(r)
	registerDiskStalledDetection(r)
	registerDjango(r)
	registerDrain(r)
	registerDrop(r)
	registerEncryption(r)
	registerFailover(r)
	registerFixtures(r)
	registerFollowerReads(r)
	registerGORM(r)
	registerGopg(r)
	registerGossip(r)
	registerHibernate(r, hibernateOpts)
	registerHibernate(r, hibernateSpatialOpts)
	registerHotSpotSplits(r)
	registerImportCancellation(r)
	registerImportDecommissioned(r)
	registerImportMixedVersions(r)
	registerImportNodeShutdown(r)
	registerImportTPCC(r)
	registerImportTPCH(r)
	registerInconsistency(r)
	registerIndexes(r)
	registerJasyncSQL(r)
	registerJepsen(r)
	registerJobsMixedVersions(r)
	registerKV(r)
	registerKVBench(r)
	registerKVContention(r)
	registerKVGracefulDraining(r)
	registerKVQuiescenceDead(r)
	registerKVRangeLookups(r)
	registerKVScalability(r)
	registerKVSplits(r)
	registerKVRestartImpact(r)
	registerKnex(r)
	registerLOQRecovery(r)
	registerLargeRange(r)
	registerLeasePreferences(r)
	registerLedger(r)
	registerLibPQ(r)
	registerLiquibase(r)
	registerLoadSplits(r)
	registerMVCCGC(r)
	registerMultiTenantDistSQL(r)
	registerMultiTenantTPCH(r)
	registerMultiTenantUpgrade(r)
	registerMultiTenantSharedProcess(r)
	registerNetwork(r)
	registerNodeJSPostgres(r)
	registerNpgsql(r)
	registerPebbleWriteThroughput(r)
	registerPebbleYCSB(r)
	registerPgjdbc(r)
	registerPgx(r)
	registerPop(r)
	registerProcessLock(r)
	registerPsycopg(r)
	registerQueue(r)
	registerQuitTransfersLeases(r)
	registerRebalanceLoad(r)
	registerReplicaGC(r)
	registerRestart(r)
	registerRestore(r)
	registerRestoreNodeShutdown(r)
	registerRoachmart(r)
	registerRoachtest(r)
	registerRubyPG(r)
	registerRustPostgres(r)
	registerSQLAlchemy(r)
	registerSQLSmith(r)
	registerSchemaChangeBulkIngest(r)
	registerSchemaChangeDuringKV(r)
	registerSchemaChangeDuringTPCC800(r)
	registerSchemaChangeIndexTPCC100(r)
	registerSchemaChangeIndexTPCC800(r)
	registerSchemaChangeInvertedIndex(r)
	registerSchemaChangeMixedVersions(r)
	registerDeclSchemaChangeCompatMixedVersions(r)
	registerSchemaChangeRandomLoad(r)
	registerScrubAllChecksTPCC(r)
	registerScrubIndexOnlyTPCC(r)
	registerSecondaryIndexesMultiVersionCluster(r)
	registerSecure(r)
	registerSequelize(r)
	registerSlowDrain(r)
	registerSyncTest(r)
	registerSysbench(r)
	registerTLP(r)
	registerTPCC(r)
	registerTPCDSVec(r)
	registerTPCE(r)
	registerTPCHBench(r)
	registerTPCHConcurrency(r)
	registerTPCHVec(r)
	registerTypeORM(r)
	registerUnoptimizedQueryOracle(r)
	registerVersion(r)
	registerYCSB(r)
	registerDeclarativeSchemaChangerJobCompatibilityInMixedVersion(r)
	registerTenantSpanStatsMixedVersion(r)
}
