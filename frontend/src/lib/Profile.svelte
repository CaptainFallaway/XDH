<script lang="ts">
  import {
    Button,
    buttonVariants
  } from "$lib/components/ui/button/index.js";
  import * as Dialog from "$lib/components/ui/dialog/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  import { SquarePen } from "lucide-svelte"
  import DatePicker from "$lib/DatePicker.svelte";
</script>
 
<Dialog.Root>
  <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
    <SquarePen />
  </Dialog.Trigger>
  <Dialog.Content class="sm:max-w-[425px]">
    <Dialog.Header>
      <Dialog.Title>Mätförrätar profil</Dialog.Title>
      <Dialog.Description>
        Här gör du ändringar till din profil och trycker sen spara.
      </Dialog.Description>
    </Dialog.Header>
    <div class="grid gap-4 py-4">
      <div class="grid grid-cols items-center gap-4">
        <Label for="name" class="text-right">Name</Label>
        <div class="col-span-3">
            <DatePicker/>
        </div>
      </div>
      <div class="grid grid-cols-4 items-center gap-4">
        <Label for="name" class="text-right">Name</Label>
        <Input id="name" class="col-span-3" />
      </div>
      <div class="grid grid-cols-4 items-center gap-4">
        <Label for="username" class="text-right">Username</Label>
        <Input id="username" class="col-span-3" />
      </div>
    </div>
    <Dialog.Footer>
      <Button type="submit">Spara</Button>
    </Dialog.Footer>
  </Dialog.Content>
</Dialog.Root>
