<script lang="ts">
  import { slide } from 'svelte/transition';
  import { ChevronDown, ChevronUp, AlertTriangle, FileDown } from 'lucide-svelte';
  import MetalSummary from './MetalSummary.svelte';
  import ErrorNotes from './ErrorNotes.svelte';
  import ScansTable from './ScansTable.svelte';
  import InvalidScansTable from './InvalidScansTable.svelte';

  import type { internal } from "../wailsjs/go/models";
  
  // Props
let {
  id,
  grouping
}: {
  id: string | null | undefined,
  grouping: internal.Grouping
} = $props();
  
  // State
  let expanded = $state(false);
  let enabledScans: Record<number, boolean> = $state({});
  
  // Initialize all scans as enabled
  // $: {
  //   if (grouping && grouping.scans) {
  //     grouping.scans.forEach((_, index) => {
  //       if (enabledScans[index] === undefined) {
  //         enabledScans[index] = true;
  //       }
  //     });
  //   }
  // }
  
  // Format date from timestamp
  function formatDate(timestamp: number): string {
    return new Date(timestamp).toLocaleDateString();
  }
  
  // Export data callback
  function exportData(): void {
    // Get only enabled scans
    const enabledScansList = grouping.scans.filter((_, index) => enabledScans[index]);
    
    // Simple callback that prints information about what would be exported
    console.log(`Exporting data for Boat ID: ${grouping.boatID}`);
    console.log(`Date Range: ${formatDate(grouping.firstDate)} - ${formatDate(grouping.lastDate)}`);
    console.log(`Number of enabled scans: ${enabledScansList.length}`);
    console.log(`Violations: ${JSON.stringify(grouping.violations)}`);
    
    // This is where you would call your Wails backend function to handle the export
    // For example: window.go.main.App.ExportGroupingData(grouping.boatID, enabledScansList);
    
    alert(`Export requested for Boat ID: ${grouping.boatID} with ${enabledScansList.length} scans`);
  }
  
  // Calculate violation counts for metals
  function getViolationCounts() {
    const counts = {
      pb: 0,
      zn: 0,
      cu: 0,
      sn: 0
    };
    
    grouping.scans.forEach(scan => {
      if (scan.violations['pb']) counts.pb++;
      if (scan.violations['zn']) counts.zn++;
      if (scan.violations['cu']) counts.cu++;
      if (scan.violations['sn']) counts.sn++;
    });
    
    return counts;
  }
  
  const violationCounts = getViolationCounts();
  
  // Handle scan toggle from child component
  function handleToggleScan(event: CustomEvent<{index: number}>) {
    const { index } = event.detail;
    enabledScans[index] = !enabledScans[index];
    enabledScans = {...enabledScans}; // Trigger reactivity
  }
  
  // Handle scan reordering from child component
  function handleReorderScans(event: CustomEvent<{scans: internal.Scan[]}>) {
    grouping.scans = event.detail.scans;
  }
</script>

<div class="card w-full bg-base-100 shadow-xl mb-4">
  <!-- Card header - always visible -->
  <div class="card-body p-4">
    <div class="flex justify-between items-center">
      <div class="flex-1">
        <h2 class="card-title text-lg">
          Boat ID: {grouping.boatID}
          {#if Object.values(grouping.violations).some(count => count > 0)}
            <span class="badge badge-error gap-1">
              <AlertTriangle size={14} />
              Violations
            </span>
          {/if}
        </h2>
        <div class="flex flex-wrap gap-2 mt-1">
          <span class="badge badge-outline">{formatDate(grouping.firstDate)} - {formatDate(grouping.lastDate)}</span>
          <span class="badge badge-outline">Unit: {grouping.unit}</span>
          <span class="badge badge-outline">Operators: {grouping.operators.join(', ')}</span>
          <span class="badge badge-outline">Scans: {grouping.scans.length}</span>
        </div>
      </div>
      
      <button 
        class="btn btn-sm btn-circle" 
        onclick={() => expanded = !expanded}
        aria-label={expanded ? "Collapse" : "Expand"}
      >
        {#if expanded}
          <ChevronUp size={18} />
        {:else}
          <ChevronDown size={18} />
        {/if}
      </button>
    </div>
  </div>
  
  <!-- Expanded content -->
  {#if expanded}
    <div transition:slide={{ duration: 300 }} class="px-4 pb-4">
      <!-- Metal readings summary -->
      <MetalSummary {violationCounts} />
      
      <!-- Error notes if any -->
      {#if grouping.errorNotes.length > 0}
        <ErrorNotes notes={grouping.errorNotes} />
      {/if}
      
      <!-- Export button -->
      <div class="flex justify-end mb-4">
        <button class="btn btn-primary" onclick={exportData}>
          <FileDown size={18} />
          Export Data
        </button>
      </div>
      
      <!-- Scans table -->
      <ScansTable 
        scans={grouping.scans} 
        unit={grouping.unit} 
        enabledScans={enabledScans}
      />
      
      <!-- Invalid scans section if any -->
      {#if grouping.invalidScans.length > 0}
        <div class="mt-6">
          <h3 class="text-lg font-semibold mb-2">Invalid Scans</h3>
          <InvalidScansTable scans={grouping.invalidScans} unit={grouping.unit} />
        </div>
      {/if}
    </div>
  {/if}
</div>