export default {
    parseUrlAndContent(description: string) {
        const urlRegex = /(http[s]?:\/\/\S+)/g;
        const matches = description.match(urlRegex);

        if (!matches) {
            return [{ type: 'text', content: description }];
        }

        const parts = description.split(urlRegex);

        return parts.reduce<Array<{ type: string; content: string; href?: string }>>(
            (result, part, index) => {
                if (index % 2 === 0) {
                    result.push({ type: 'text', content: part });
                } else {
                    result.push({ type: 'link', content: part, href: part });
                }
                return result;
            }, []);
    }
};