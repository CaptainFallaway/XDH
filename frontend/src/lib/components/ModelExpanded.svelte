<script lang="ts">
    import { slide, fade } from "svelte/transition";
    import * as Tabs from "$lib/components/ui/tabs";
    import Table from "$lib/components/Table.svelte";
    import { cubicOut } from "svelte/easing";
    import Info from "$lib/components/Info.svelte";
    import { toggleValue } from "$lib/globalstores.ts";
    import { internal } from "../wailsjs/go/models.ts";

    export let model: internal.Grouping

    const expandMotion = {
        duration: 500,
        easing: cubicOut,
    };

    let tabValue = "Valid";
</script>

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
