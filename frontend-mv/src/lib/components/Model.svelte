<script lang="ts">
    import { fly, fade } from "svelte/transition";

    import { toggleValue } from "$lib/globalstores.ts";
    // import { Badge } from "$lib/components/ui/badge/";
    import { Separator } from "$lib/components/ui/separator";
    import Info from "$lib/components/Info.svelte";
    import { ChevronUp, ChevronDown } from "lucide-svelte";
    import ModelExpanded from "./ModelExpanded.svelte";
    import { internal } from "../wailsjs/go/models.ts";

    export let id = "0";
    export let model: internal.Grouping;

    // let status;
    let statusColor: string;

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
        <Info top="Båt ID" bottom={model.boatID} />
        <Info top="Mätförättare" bottom={model.operators.join(", ")} />
        <Info top="Förta Tid" bottom={model.firstDate.text} />
        <Info top="Sista Tid" bottom={model.lastDate.text} />
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
        <ModelExpanded {model}/>
    {/if}
</div>
