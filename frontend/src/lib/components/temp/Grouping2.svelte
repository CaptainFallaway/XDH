<script lang="ts">
import {
  slide
} from 'svelte/transition';
import {
  flip
} from 'svelte/animate';
import type {
  internal
} from "../../wailsjs/go/models";
import {
  FileDown,
  ChevronDown,
  ChevronUp,
  AlertTriangle
} from 'lucide-svelte';
// import { jsPDF } from 'jspdf';
// import 'jspdf-autotable';
import {
  onMount
} from 'svelte';

// // Define types based on the Go structs
// type Scan = {
//   reading: number;
//   duration: number;
//   operator: string;
//   date: number;
//   pb: number;
//   zn: number;
//   cu: number;
//   sn: number;
//   violations: Record<string, boolean>;
// };

// type Grouping = {
//   index: number;
//   boatID: string;
//   firstDate: number;
//   lastDate: number;
//   unit: string;
//   scans: Scan[];
//   invalidScans: Scan[];
//   errorNotes: string[];
//   violations: Record<string, number>;
//   operators: string[];
// };

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
let enabledScans: Record < number, boolean > = {};

// Initialize all scans as enabled
onMount(() => {
  grouping.scans.forEach((_, index) => {
    enabledScans[index] = true;
  });
});

// Format date from timestamp
function formatDate(timestamp: number): string {
  return new Date(timestamp*1000).toLocaleDateString();
}

// Format time from timestamp
function formatTime(timestamp: number): string {
  return new Date(timestamp*1000).toLocaleTimeString();
}

// Toggle scan enabled state
function toggleScan(index: number): void {
  enabledScans[index] = !enabledScans[index];
}

// Move scan up in the list
function moveScanUp(index: number): void {
  if (index > 0) {
    const newScans = [...grouping.scans];
    [newScans[index - 1], newScans[index]] = [newScans[index], newScans[index - 1]];
    grouping.scans = newScans;
  }
}

// Move scan down in the list
function moveScanDown(index: number): void {
  if (index < grouping.scans.length - 1) {
    const newScans = [...grouping.scans];
    [newScans[index], newScans[index + 1]] = [newScans[index + 1], newScans[index]];
    grouping.scans = newScans;
  }
}

// Export to PDF
function exportToPDF(): void {
  // const doc = new jsPDF();

  // // Add title
  // doc.setFontSize(18);
  // doc.text(`Boat ID: ${grouping.boatID}`, 14, 22);

  // // Add metadata
  // doc.setFontSize(12);
  // doc.text(`Date Range: ${formatDate(grouping.firstDate)} - ${formatDate(grouping.lastDate)}`, 14, 30);
  // doc.text(`Unit: ${grouping.unit}`, 14, 36);
  // doc.text(`Operators: ${grouping.operators.join(', ')}`, 14, 42);

  // // Add violations summary if any
  // if (Object.keys(grouping.violations).length > 0) {
  //   doc.text('Violations:', 14, 48);
  //   let yPos = 54;
  //   Object.entries(grouping.violations).forEach(([metal, count]) => {
  //     doc.text(`- ${metal}: ${count}`, 20, yPos);
  //     yPos += 6;
  //   });
  // }

  // // Add scans table (only enabled scans)
  // const enabledScansList = grouping.scans.filter((_, index) => enabledScans[index]);

  // if (enabledScansList.length > 0) {
  //   const tableData = enabledScansList.map(scan => [
  //     scan.reading.toString(),
  //     scan.duration.toFixed(2),
  //     scan.operator,
  //     formatDate(scan.date),
  //     scan.pb.toFixed(2),
  //     scan.zn.toFixed(2),
  //     scan.cu.toFixed(2),
  //     scan.sn.toFixed(2),
  //     Object.keys(scan.violations).filter(key => scan.violations[key]).join(', ')
  //   ]);

  //   // @ts-ignore - jspdf-autotable types not available
  //   doc.autoTable({
  //     startY: Object.keys(grouping.violations).length > 0 ? 60 + Object.keys(grouping.violations).length * 6 : 60,
  //     head: [['Reading', 'Duration', 'Operator', 'Date', 'Pb', 'Zn', 'Cu', 'Sn', 'Violations']],
  //     body: tableData,
  //   });
  // }

  // // Add error notes if any
  // if (grouping.errorNotes.length > 0) {
  //   doc.addPage();
  //   doc.text('Error Notes:', 14, 20);
  //   let yPos = 26;
  //   grouping.errorNotes.forEach(note => {
  //     doc.text(`- ${note}`, 20, yPos);
  //     yPos += 6;
  //   });
  // }

  // // Save the PDF
  // doc.save(`boat-${grouping.boatID}-report.pdf`);
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

// Get color class based on violation status
function getViolationColorClass(value: number, threshold: number): string {
  return value > threshold ? 'text-error' : 'text-success';
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
        <div class="stats shadow w-full mb-4">
            <div class="stat">
                <div class="stat-title">Lead (Pb)</div>
                <div class="stat-value text-2xl" class:text-error={violationCounts.pb > 0}>
                    {violationCounts.pb > 0 ? violationCounts.pb : 'OK'}
                </div>
                <div class="stat-desc">Violations</div>
            </div>

            <div class="stat">
                <div class="stat-title">Zinc (Zn)</div>
                <div class="stat-value text-2xl" class:text-error={violationCounts.zn > 0}>
                    {violationCounts.zn > 0 ? violationCounts.zn : 'OK'}
                </div>
                <div class="stat-desc">Violations</div>
            </div>

            <div class="stat">
                <div class="stat-title">Copper (Cu)</div>
                <div class="stat-value text-2xl" class:text-error={violationCounts.cu > 0}>
                    {violationCounts.cu > 0 ? violationCounts.cu : 'OK'}
                </div>
                <div class="stat-desc">Violations</div>
            </div>

            <div class="stat">
                <div class="stat-title">Tin (Sn)</div>
                <div class="stat-value text-2xl" class:text-error={violationCounts.sn > 0}>
                    {violationCounts.sn > 0 ? violationCounts.sn : 'OK'}
                </div>
                <div class="stat-desc">Violations</div>
            </div>
        </div>

        <!-- Error notes if any -->
        {#if grouping.errorNotes.length > 0}
        <div class="alert alert-warning mb-4">
            <AlertTriangle size={18} />
            <div>
                <h3 class="font-bold">Notes</h3>
                <ul class="list-disc list-inside">
                    {#each grouping.errorNotes as note}
                    <li>{note}</li>
                    {/each}
                </ul>
            </div>
        </div>
        {/if}

        <!-- Export button -->
        <div class="flex justify-end mb-4">
            <button class="btn btn-primary" onclick={exportToPDF}>
                <FileDown size={18} />
                Export to PDF
            </button>
        </div>

        <!-- Scans table -->
        <div class="overflow-x-auto">
            <table class="table table-zebra w-full">
                <thead>
                    <tr>
                        <th>Actions</th>
                        <th>Reading</th>
                        <th>Duration</th>
                        <th>Operator</th>
                        <th>Date/Time</th>
                        <th>Pb ({grouping.unit})</th>
                        <th>Zn ({grouping.unit})</th>
                        <th>Cu ({grouping.unit})</th>
                        <th>Sn ({grouping.unit})</th>
                        <th>Enabled</th>
                    </tr>
                </thead>
                <tbody>
                    {#each grouping.scans as scan, i (scan.reading)}
                    <tr
                        animate:flip={{ duration: 300 }}
                        class:opacity-50={!enabledScans[i]}
                        >
                        <td class="flex gap-1">
                            <button
                                class="btn btn-xs"
                                onclick={() => moveScanUp(i)}
                                disabled={i === 0}
                                >
                                ↑
                            </button>
                            <button
                                class="btn btn-xs"
                                onclick={() => moveScanDown(i)}
                                disabled={i === grouping.scans.length - 1}
                                >
                                ↓
                            </button>
                        </td>
                        <td>{scan.reading}</td>
                        <td>{scan.duration.toFixed(2)}</td>
                        <td>{scan.operator}</td>
                        <td>
                            {formatDate(scan.date)}
                            <br />
                            <span class="text-xs opacity-70">{formatTime(scan.date)}</span>
                        </td>
                        <td class={scan.violations['pb'] ? 'text-error font-bold' : ''}>{scan.pb.toFixed(2)}</td>
                        <td class={scan.violations['zn'] ? 'text-error font-bold' : ''}>{scan.zn.toFixed(2)}</td>
                        <td class={scan.violations['cu'] ? 'text-error font-bold' : ''}>{scan.cu.toFixed(2)}</td>
                        <td class={scan.violations['sn'] ? 'text-error font-bold' : ''}>{scan.sn.toFixed(2)}</td>
                        <td>
                            <input
                                type="checkbox"
                                class="toggle toggle-primary toggle-sm"
                                checked={enabledScans[i]}
                                onchange={() => toggleScan(i)}
                            />
                        </td>
                    </tr>
                    {/each}
                </tbody>
            </table>
        </div>

        <!-- Invalid scans section if any -->
        {#if grouping.invalidScans.length > 0}
        <div class="mt-6">
            <h3 class="text-lg font-semibold mb-2">Invalid Scans</h3>
            <div class="overflow-x-auto">
                <table class="table table-zebra w-full">
                    <thead>
                        <tr>
                            <th>Reading</th>
                            <th>Duration</th>
                            <th>Operator</th>
                            <th>Date/Time</th>
                            <th>Pb ({grouping.unit})</th>
                            <th>Zn ({grouping.unit})</th>
                            <th>Cu ({grouping.unit})</th>
                            <th>Sn ({grouping.unit})</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each grouping.invalidScans as scan}
                        <tr class="opacity-60">
                            <td>{scan.reading}</td>
                            <td class="text-error">{scan.duration.toFixed(2)}</td>
                            <td>{scan.operator}</td>
                            <td>
                                {formatDate(scan.date)}
                                <br />
                                <span class="text-xs opacity-70">{formatTime(scan.date)}</span>
                            </td>
                            <td>{scan.pb.toFixed(2)}</td>
                            <td>{scan.zn.toFixed(2)}</td>
                            <td>{scan.cu.toFixed(2)}</td>
                            <td>{scan.sn.toFixed(2)}</td>
                        </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
        </div>
        {/if}
    </div>
    {/if}
</div>
