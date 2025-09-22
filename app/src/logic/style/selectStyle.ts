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
            color: selectStyle.isSelected(value) ? "#4d5161" : "#AEB7BD",
        }
    },
}