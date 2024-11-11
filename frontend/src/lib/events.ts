import * as wailsrt from "$lib/wailsjs/runtime/runtime"
import { toast } from "svelte-sonner";

wailsrt.EventsOn("error", () => {
    toast.error("An error occurred");
})