<script lang="ts">
    import { flip } from 'svelte/animate';
  
    let rows = $state([
      { id: 1, name: 'John Doe', email: 'john@example.com' },
      { id: 2, name: 'Jane Smith', email: 'jane@example.com' },
      { id: 3, name: 'Bob Johnson', email: 'bob@example.com' },
      { id: 4, name: 'Alice Brown', email: 'alice@example.com' },
    ]);
  
    function moveRow(index: number, direction: number) {
      if ((direction === -1 && index > 0) || (direction === 1 && index < rows.length - 1)) {
        const newRows = [...rows];
        const temp = newRows[index];
        newRows[index] = newRows[index + direction];
        newRows[index + direction] = temp;
        rows = newRows;
      }
    }
  </script>
  
  <div class="overflow-x-auto">
    <table class="table">
      <thead>
        <tr>
          <th>Actions</th>
          <th>Name</th>
          <th>Email</th>
        </tr>
      </thead>
      <tbody>
        {#each rows as row, index (row.id)}
          <tr animate:flip={{ duration: 300 }}>
            <td>
              <div class="join">
                <button
                  class="btn btn-sm join-item"
                  onclick={() => moveRow(index, -1)}
                  disabled={index === 0}
                >
                  ↑
                </button>
                <button
                  class="btn btn-sm join-item"
                  onclick={() => moveRow(index, 1)}
                  disabled={index === rows.length - 1}
                >
                  ↓
                </button>
              </div>
            </td>
            <td>{row.name}</td>
            <td>{row.email}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>