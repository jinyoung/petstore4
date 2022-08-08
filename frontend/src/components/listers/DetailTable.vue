<template>

    <v-data-table
            :headers="headers"
            :items="detail"
            :items-per-page="5"
            class="elevation-1"
    ></v-data-table>

</template>

<script>

    export default {
        name: 'DetailList',
        props: {
            value: Object,
            editMode: Boolean,
            isNew: Boolean
        },
        data: () => ({
            headers: [
            ],
            detail : [],
        }),
        async created() {
            if(!this.value) this.value = [];
        },
        methods: {
        }
    }
</script>

