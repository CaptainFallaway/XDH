<script lang="ts">
    export let top = "";
    export let bottom: any;
    export let classes = "";

    if (bottom == null) {
        bottom = "None";
    }
</script>

<div class="{classes} justify-start flex flex-col whitespace-nowrap">
    <span class="top">{top}</span>
    <span>{bottom}</span>
</div>

<style>
    .top {
        opacity: 0.7;
        font-size: small;
    }
</style>
