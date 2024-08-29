<script>
    import * as Table from "$lib/components/ui/table/index.js";

    export let scans = [];

    if (scans == null) {
        scans = [];
    }
</script>

<Table.Root>
    <Table.Header>
        <Table.Row>
            <Table.Head>#</Table.Head>
            <Table.Head>Reading No</Table.Head>
            <Table.Head>Duration</Table.Head>
            <Table.Head>Time</Table.Head>
            <Table.Head>Pb</Table.Head>
            <Table.Head>Sn</Table.Head>
        </Table.Row>
    </Table.Header>
    <Table.Body>
        {#each scans as scan, i}
            <Table.Row>
                <Table.Cell>{i + 1}</Table.Cell>
                <Table.Cell>{scan.reading}</Table.Cell>
                <Table.Cell>{scan.duration}</Table.Cell>
                <Table.Cell>{scan.time}</Table.Cell>
                <Table.Cell>{scan.pb.value}</Table.Cell>
                <Table.Cell>{scan.sn.value}</Table.Cell>
            </Table.Row>
        {/each}
    </Table.Body>
</Table.Root>
