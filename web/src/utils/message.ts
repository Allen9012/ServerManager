import { ElMessage } from 'element-plus';

let messageDom: any = null;
const Message: any = (options) => {
    if (messageDom) messageDom.close();
    messageDom = ElMessage(options);
};

export const MsgSuccess = (message) => {
    Message.success({
        message: message,
        type: 'success',
        showClose: true,
        duration: 3000,
    });
};

export const MsgError = (message) => {
    Message.error({
        message: message,
        type: 'error',
        showClose: true,
        duration: 3000,
    });
};

export const MsgWarning = (message) => {
    Message.warning({
        message: message,
        type: 'warning',
        showClose: true,
        duration: 3000,
    });
};
