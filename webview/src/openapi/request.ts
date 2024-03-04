export async function httpRequest(input: string, options: RequestInit = {}) {
    const headers = { 'Content-Type': 'application/json' };
    options.headers = Object.assign(headers, options.headers);

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
            throw data.Error;
        }
        return data.Payload;
    } catch (error) {
        window.postMessage({ message: String(error), type: 'danger' });
        throw error;
    }
}
