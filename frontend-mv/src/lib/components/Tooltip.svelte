<script lang="ts">
    import * as Tooltip from "$lib/components/ui/tooltip"
    import { Button } from "$lib/components/ui/button";

    export let tip: string = "";
</script>

<Tooltip.Root>
  <Tooltip.Trigger>
    <slot>
        <Button>REPLACE ME</Button>
    </slot>
  </Tooltip.Trigger>
  <Tooltip.Content>
    <p>{tip}</p>
  </Tooltip.Content>
</Tooltip.Root>
