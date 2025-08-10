export const header = {
    getColor: function(menu, trueMenu){
        if (menu === trueMenu) {
            return '#FFFFFF';
        } else {
            return '#739DDB';
        }
    },

    fontsize: '24px',
    getLinkStyle: function(menu, trueMenu){
        const weight = menu === trueMenu ? 'bold' : 'normal';
        return {
            color: header.getColor(menu,trueMenu),
            fontWeight: weight,
            fontSize: header.fontsize
        }
    },

    getUnderlineStyle: function(menu, trueMenu) {
        return {
            width: '100%',
            height: '8px',
            backgroundColor: menu === trueMenu ? '#66BEBC' : '',
        };
    },

    getIconStyle: function(trueMenu) {
        return {
            // color: header.getColor(trueMenu),
            width: '30px',
            height: '30px',
            display: 'inline-block',
            position: 'relative',
            top: '5px',
        }
    },

    log: "StudyTime",
    subject: "AddSubject",

    list: "TodoList",
    studylog: "StudyLog",
    track: "TimeTrack",
    letter: "Note",
}