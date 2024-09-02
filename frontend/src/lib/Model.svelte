<script lang="ts">
    import { cubicOut } from "svelte/easing";
    import { fly, slide, fade } from "svelte/transition";

    import { toggleValue } from "$lib/stores.ts";
    // import { Badge } from "$lib/components/ui/badge/";
    import { Separator } from "$lib/components/ui/separator";
    import Info from "$lib/Info.svelte";
    import * as Tabs from "$lib/components/ui/tabs/index.ts";
    import Table from "$lib/Table.svelte";
    import { ChevronUp, ChevronDown } from "lucide-svelte";
    import { data_pipeline } from "./wailsjs/go/models";

    export let id = "0";
    export let model: data_pipeline.Grouping;

    const expandMotion = {
        duration: 500,
        easing: cubicOut,
    };

    // let status;
    let statusColor;

    switch (model.violations[$toggleValue]) {
        case 0:
            // status = "safe";
            statusColor = "green";
            break;
        case 1:
            // status = "warning";
            statusColor = "yellow";
            break;
        default:
            // status = "danger";
            statusColor = "red";
            break;
    }

    let show = false;

    function toggleShow() {
        show = !show;
    }

    function handleKeyPress(event: KeyboardEvent) {
        if (event.key === "Enter" || event.key === " ") {
            toggleShow();
        }
    }

    let tabValue = "Valid";
</script>

<!-- shadow-red-500 shadow-yellow-500 shadow-green-500 (Not used right now) -->
<!-- border-l-red-500 border-l-yellow-500 border-l-green-500 -->

<div
    {id}
    role="button"
    on:click={toggleShow}
    on:keydown={handleKeyPress}
    tabindex="0"
    in:fly|global={{ delay: 200, y: 200, duration: 200 }}
    out:fade|global={{ duration: 200 }}
    class="border-2 rounded-lg right-0 left-0 m-5 content-center shadow-xl border-l-4 border-l-{statusColor}-500"
>
    <div aria-label="info" class="flex gap-10 m-8 content-center">
        <Info top="Boat ID" bottom={model.boatID} />
        <Info top="Operators" bottom={model.operators.join(", ")} />
        <Info top="First Date" bottom={model.firstDate} />
        <Info top="Last Date" bottom={model.lastDate} />
        <div class="w-full flex items-center justify-end">
            {#if show}
                <div in:fade|global>
                    <ChevronDown size={48} />
                </div>
            {:else}
                <div in:fade|global>
                    <ChevronUp size={48} />
                </div>
            {/if}
        </div>
    </div>
    {#if show}
        <Separator />
        <div
            class="m-2 flex cursor-default"
            transition:slide|global={{ ...expandMotion }}
            on:click|stopPropagation
            on:keydown|stopPropagation
            role="button"
            tabindex="0"
        >
            <div class="w-full mx-5 mb-5 flex flex-col" transition:fade|global>
                <div class="w-full flex content-center justify-center gap-8">
                    <Info
                        top="Violations"
                        bottom={model.violations[$toggleValue]}
                    />
                    <Info
                        top="Scans"
                        bottom={model.scans != null ? model.scans.length : 0}
                    />
                    <Info
                        top="Invalid Scans"
                        bottom={model.invalidScans != null
                            ? model.invalidScans.length
                            : 0}
                    />
                    <Info top="Notes" bottom={model.errorNotes} />
                </div>
                <Tabs.Root bind:value={tabValue} class="w-full">
                    <Tabs.List>
                        <Tabs.Trigger value="Valid">Valid</Tabs.Trigger>
                        <Tabs.Trigger value="Invalid">Invalid</Tabs.Trigger>
                    </Tabs.List>
                    <Tabs.Content value="Valid">
                        {#key tabValue}
                            <div in:fade|global>
                                <Table scans={model.scans} />
                            </div>
                        {/key}
                    </Tabs.Content>
                    <Tabs.Content value="Invalid">
                        {#key tabValue}
                            <div in:fade|global>
                                <Table scans={model.invalidScans} />
                            </div>
                        {/key}
                    </Tabs.Content>
                </Tabs.Root>
            </div>
        </div>
    {/if}
</div>
