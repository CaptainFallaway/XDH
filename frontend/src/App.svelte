<script lang="ts">
    import { NewSession, ListSessions, NewSessionData, OpenFileDialog, GetGroupings, SetSession } from "./lib/wailsjs/go/app/App";
    import type { internal } from "./lib/wailsjs/go/models";

    let groupings: internal.Grouping[] = $state([]);

    $inspect(groupings);

    async function newSession() {
        let sessions = await ListSessions();
        let session = sessions[0];
        await SetSession(session.uid);
        groupings = await GetGroupings("Pb"); 
    }
</script>

<div class="flex justify-center m-auto flex-col">
    <div class="flex flex-col w-[50vw] m-auto space-y-4 mb-4">
        <button onclick={newSession}>Load Session</button>
        <button onclick={async () => console.log(await ListSessions())}>List Sessions</button>
    </div>
    {#each groupings as group (group.boatID)}
       <textarea class="w-[50vw] h-[100vh] m-auto mb-4 border border-black rounded" name="text" id={group.boatID}>
            {JSON.stringify(group, null, 2)}
       </textarea> 
    {/each}
</div>

<style>
    button {
        background-color: black;
        border-radius: 0.5rem;
        color: white;
        height: 3em;
        width: 50%;
        margin: auto;
        transition: all 0.3s;
    }

    button:hover {
        background-color: grey;
    }
</style>