<script>
    import { cubicOut } from "svelte/easing";
    import { fly, slide } from "svelte/transition";

    import { toggleValue } from "$lib/stores.js";
    import { Badge } from "$lib/components/ui/badge/";
    import Info from "$lib/Info.svelte";

    export let id = '0';
    export let model = {};

    const expandMotion = {
        duration: 500,
        easing: cubicOut,
    }

    let status;
    let statusColor;

    switch (model.violations[$toggleValue]) {
        case 0:
            status = "safe"
            statusColor = "green"
            break;
        case 1:
            status = "warning"
            statusColor = "yellow"
            break;
        default:
            status = "danger"
            statusColor = "red"
            break;
    }

    let show = false;

    function toggleShow() {
        show = !show;
    }

    function handleKeyPress(event) {
        if (event.key === "Enter" || event.key === " ") {
            toggleShow();
        }
    }
</script>

<div {id} role="button" on:click={toggleShow} on:keydown={handleKeyPress} tabindex="0" transition:fly={{y: 200, duration: 500}} class="border-2 rounded-lg right-0 left-0 m-3">
    <div aria-label="info" class="flex gap-4 m-2 content-center">
        <div class="content-center">
            <Badge {id} style="background: {statusColor}">
                <div class="w-8 h-5"/>
            </Badge>
        </div>
        <Info desc="Index" info={model.index} />
    </div>
    {#if show}
        <div class="m-2" in:slide={{...expandMotion}} out:slide={{...expandMotion}}>
            <div class="w-full h-[100px] rounded-lg border-2"></div>
        </div>
    {/if}
</div>