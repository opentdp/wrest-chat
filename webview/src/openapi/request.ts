export async function httpRequest(input: string, options: RequestInit = {}) {
    options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);

    try {
        const response = await fetch(input, options);
        if (response.status < 200 || response.status > 300) {
            throw new Error(response.statusText);
        }
        const data = await response.json();
        if (data.Message) {
            window.postMessage(data.Message);
        }
        if (data.Error) {
            throw data.Error;
        }
        return data.Payload;
    } catch (error) {
        const message = String(error);
        window.postMessage({ message, classname: 'bg-danger text-light' });
        throw error;
    }

}
