<script lang="ts">
    import * as Table from "$lib/components/ui/table";
    import { internal } from "$lib/wailsjs/go/models.ts";

    export let scans: Array<internal.ScanRow>;

    if (scans == null) {
        scans = [];
    }
</script>

<Table.Root>
    <Table.Header>
        <Table.Row>
            <Table.Head>#</Table.Head>
            <Table.Head>Scan Nummer</Table.Head>
            <Table.Head>Duration</Table.Head>
            <Table.Head>Tid</Table.Head>
            <Table.Head>Tenn</Table.Head>
            <Table.Head>Bly</Table.Head>
        </Table.Row>
    </Table.Header>
    <Table.Body>
        {#each scans as scan, i}
            <Table.Row>
                <Table.Cell>{i + 1}</Table.Cell>
                <Table.Cell>{scan.reading}</Table.Cell>
                <Table.Cell>{scan.duration}</Table.Cell>
                <Table.Cell>{scan.time.text}</Table.Cell>
                <Table.Cell>{scan.sn.value}</Table.Cell>
                <Table.Cell>{scan.pb.value}</Table.Cell>
            </Table.Row>
        {/each}
    </Table.Body>
</Table.Root>
