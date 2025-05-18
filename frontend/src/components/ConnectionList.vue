<template>
    <main>
        <dev v-for="item in list">{{ item }}</dev>
    </main>
</template>

<script setup>
    import {ref} from "vue"
    let list = ref()
</script>


<style scoped></style>