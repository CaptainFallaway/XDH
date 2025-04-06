<script lang="ts">
    import { Metal } from './lib/Global.svelte.ts';
	import Grouping from './lib/components/Grouping.svelte';
    import * as app from "./lib/wailsjs/go/app/App";
    import type { internal } from "./lib/wailsjs/go/models";

    let sessions: internal.SessionInfo[] = $state([]);
    let groupings: internal.Grouping[] = $state([]);

    let session = $state("");
    let metal = $state(Metal);
    
    async function fetchSessions() {
        sessions = await app.ListSessions();
    }

    async function newSession() {
        let data = await app.NewSessionData();
        data.surveyor = "retep";
        data.date = new Date().getTime();
        data.location = "Lomma"
        data.instrumentSerial = "1234";

        let path = await app.OpenFileDialog();

        console.log(path);       

        session = await app.CreateSession(data, path);

        fetchSessions();
    }

    async function deleteSession() {
        if (session == "") {
            return;
        }

        await app.DeleteSession(session);

        session = "";
        groupings = [];
        fetchSessions();
    }

    $effect(() => {
        (async () => {
            await fetchSessions();
        })();
    });

    async function getGroupings(session: string, metal: string): Promise<internal.Grouping[]> {
        const temp = await app.GetGroupings(session, metal);
        groupings = temp;
        return temp;
    }

    $inspect(groupings);
    $inspect(sessions);
    $inspect(session);
    $inspect(metal);

    // document["app"] = app;
</script>

<div class="flex justify-center m-auto flex-col mx-4 space-y-2">
    <div class="flex flex-col w-[50vw] m-auto space-y-4 mb-4">
        <button class="btn" onclick={newSession}>new session</button>
        <button class="btn" onclick={deleteSession}>delete session</button>
        <button class="btn" onclick={fetchSessions}>refresh sessions</button>
        <div class="flex flex-row space-x-2 content-center justify-center">
            <select class="border border-black rounded" bind:value={session} name="sessionSelect" id="sessionSelect">
                {#each sessions as session}
                    <option value={session.uid}>{session.uid}</option>
                {/each}
            </select>
            <select class="border border-black rounded" bind:value={metal} name="metalSelect" id="metalSelect">
                <option value="Sn">Sn</option>
                <option value="Pb">Pb</option>
                <option value="Cu">Cu</option>
                <option value="Zn">Zn</option>
            </select>
        </div>
    </div>

    {#await getGroupings(session, metal)}
        <p>loading...</p>
    {:then groupings} 
    {#each groupings as grouping}
    <!-- <textarea class="w-[50vw] h-[100vh] m-auto mb-4 border border-black rounded" name="text" id={grouping.boatID}>
        {JSON.stringify(grouping, null, 2)}
        </textarea>  -->
        <Grouping id={grouping.index.toString()} {grouping} />
    {/each}
    {/await}
</div>

<!-- <style>
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
</style> -->