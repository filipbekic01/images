new Vue({
    el: "#app",
    delimiters: ["{!", "!}"],
    data: {
        omg: 'yes'
    },
    created() {
        console.log('wow');
    }
});
