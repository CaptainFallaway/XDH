<script lang="ts">
    import { fade } from 'svelte/transition';
	  import { Metal } from '../Global.svelte.ts';
    import type { internal } from "../wailsjs/go/models";
    let { id, grouping }: { id: string | null | undefined, grouping: internal.Grouping } = $props();

    let violationCount = grouping.violations[Metal];

    let ok = violationCount == 1;
    let bad = violationCount > 1;

    $inspect(grouping);
</script>

{#snippet info(title: string, desc: any)}
    <div class="flex flex-col space-y-2 p-2">
        <strong>{title}</strong>
        <p>{desc}</p>
    </div>
{/snippet}

<div {id} transition:fade class:ok class:bad class="flex flex-row content-center w-full p-4 border border-[greenYellow] rounded-md">
  {@render info("BoatID", grouping.boatID)}
  {@render info("Första", grouping.firstDate)}
  {@render info("Sista", grouping.lastDate)}
  {@render info("Violations", grouping.violations[Metal])}
</div>

<style>
  div.ok {
    border-color: yellow;
  }

  div.bad {
    border-color: red;
  }
</style>