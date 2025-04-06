<script lang="ts">
	import Expanded from './Expanded.svelte';
    import type { internal } from "../wailsjs/go/models";
    import { ChevronDown, ChevronUp } from "lucide-svelte";
    import { fade, fly } from "svelte/transition";
    import { Metal } from "../Global.svelte";
    let { id, grouping }: { id: string | null | undefined, grouping: internal.Grouping } = $props();

    let show = $state(false);
</script>

{#snippet info(title: string, desc: any)}
    <div class="justify-start flex flex-col space-y-2 p-2 whitespace-nowrap">
        <strong>{title}</strong>
        <p>{desc}</p>
    </div>
{/snippet}

<button
  {id}
  in:fly|global={{ delay: 200, y: 200, duration: 200 }}
  out:fade|global={{ duration: 200 }}
  class:border-success={grouping.violations[Metal] === 0}
  class:border-warning={grouping.violations[Metal] === 1}
  class:border-error={grouping.violations[Metal] > 1}
  class="flex content-center items-center w-full p-4 bg-base-200 border border-success border-error border-warning rounded-md"
  onclick={() => show = !show}>
  {@render info("Båt ID", grouping.boatID)}
  {@render info("Mätförättare", grouping.operators.join(", "))}
  {@render info("Första skanning", new Date(grouping.firstDate * 1000).toLocaleString())}
  {@render info("Sista skanning", new Date(grouping.lastDate * 1000).toLocaleString())}
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
</button>