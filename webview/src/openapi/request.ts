export async function httpRequest(input: string, options: RequestInit = {}) {
    const headers: Record<string, string> = {
        'Content-Type': 'application/json'
    };
    // 设置 Token
    const token = sessionStorage.getItem('token');
    if (token && token.trim().length > 1) {
        headers['Authorization'] = `Bearer ${token}`;
    }
    // 设置 Header
    options.headers = Object.assign(headers, options.headers);
    // 发起 HTTP 请求
    try {
        const response = await fetch(input, options);
        if (response.status < 200 || response.status > 300) {
            throw new Error(response.statusText);
        }
        const data = await response.json();
        if (data.Message) {
            window.postMessage({ message: data.Message, type: 'success' });
        }
        if (data.Error) {
            if (data.Error.Message) {
                throw new Error(data.Error.Message);
            }
            throw data.Error;
        }
        return data.Payload;
    } catch (error) {
        window.postMessage({ message: String(error), type: 'danger' });
        throw error;
    }
}
