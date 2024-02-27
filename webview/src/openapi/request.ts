export function httpRequest(input: string, options: RequestInit = {}) {
    options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);

    return fetch(input, options).then(response => {
        if (response.status >= 200 && response.status < 300) {
            return response.json().then(data => {
                if (data.Error) {
                    throw data.Error;
                }
                return data.Payload;
            });
        } else {
            throw response;
        }
    });
}
