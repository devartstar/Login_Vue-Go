<template>
  <div>{{ message }}</div>
</template>

<script lang="ts">
import { onMounted, ref } from "vue";
import { useStore } from "vuex";
export default {
  name: "HomeComponent",
  setup() {
    const message = ref("You are not logged in");
    const store = useStore();
    onMounted(async () => {
      try {
        const response = await fetch("http://localhost:8000/api/user", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        });
        const content = await response.json();
        message.value = `Hi ${content.name}`;
        await store.dispatch("setAuth", true);
      } catch (e) {
        await store.dispatch("setAuth", false);
      }
    });
    return {
      message,
    };
  },
};
</script>
