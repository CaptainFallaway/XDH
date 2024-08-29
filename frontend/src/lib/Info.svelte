<script>
    export let top = "";
    export let bottom = "";
</script>

<div class="justify-start flex flex-col whitespace-nowrap">
    <span class="top">{top}</span>
    <span>{bottom}</span>
</div>

<style>
    .top {
        opacity: 0.7;
        font-size: small;
    }
</style>
