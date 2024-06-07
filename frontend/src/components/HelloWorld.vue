<template>
  <div>
    <input v-model="text" placeholder="Enter text" />
    <button @click="addRecord">Add record</button>
    <ul>
      <li v-for="record in records" :key="record.id">{{ record.text }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      text: "",
      records: [],
    };
  },
  methods: {
    async addRecord() {
      const response = await fetch("/api/records", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ text: this.text }),
      });
      const record = await response.json();
      this.records.push(record);
      this.text = "";
    },
  },
  async mounted() {
    const response = await fetch("/api/records");
    this.records = await response.json();
  },
};
</script>
