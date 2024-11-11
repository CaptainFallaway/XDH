<script lang="ts">
    import Tooltip from "./Tooltip.svelte";
    import * as Tabs from "$lib/components/ui/tabs/index.ts";

    import SessionsSheet from "$lib/SessionsSheet.svelte";
    import Profile from "$lib/Profile.svelte"

    import { toggleValue } from "$lib/globalstores";
</script>

<div class="flex flex-row right-0 left-0 m-5">
    <div class="basis-1/2 text-left flex gap-1">
        <SessionsSheet /> 
        <Profile />
    </div>
    <div class="basis-1/2 flex justify-end">
        <Tooltip tip="Vilken metall som skall sorteras med">
            <Tabs.Root bind:value={$toggleValue} class="justify-start">
                <Tabs.List class="shadow-xl">
                    <Tabs.Trigger value="Sn">Tenn</Tabs.Trigger>
                    <Tabs.Trigger value="Pb">Bly</Tabs.Trigger>
                    <!-- <Tabs.Trigger value="Cu">Cu</Tabs.Trigger>
                    <Tabs.Trigger value="Zn">Zn</Tabs.Trigger> -->
                </Tabs.List>
            </Tabs.Root>
        </Tooltip>
    </div>
</div>
