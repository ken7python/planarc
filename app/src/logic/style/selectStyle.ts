export const selectStyle = {
    isSelected: function(value: string){
        if (value) {
            return value;
        } else {
            return false;
        }
    },
    getSelectStyle: function(value) {
        return {
            color: selectStyle.isSelected(value) ? "#000" : "#AEB7BD",
        }
    },
}