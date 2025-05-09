<script lang="ts">
  import { AlertTriangle } from 'lucide-svelte';
  
  // Props
  export let notes: string[];
</script>

<div class="alert alert-warning mb-4">
  <AlertTriangle size={18} />
  <div>
    <h3 class="font-bold">Notes</h3>
    <ul class="list-disc list-inside">
      {#each notes as note}
        <li>{note}</li>
      {/each}
    </ul>
  </div>
</div>