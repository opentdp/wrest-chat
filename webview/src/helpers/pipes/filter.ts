import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
    name: 'filter'
})
export class FilterPipe implements PipeTransform {
    transform<T>(input: T[], filterBy: string, filterValue: string): T[] {
        if (!input || !filterBy || !filterValue) {
            return input;
        }

        const keys = filterBy.toLowerCase().split(',');
        const value = filterValue.toLowerCase();

        return input.filter(item => {
            return keys.some(k => {
                if (typeof item == 'object' && item && k in item) {
                    const obj = item as { [key: string]: unknown };
                    return String(obj[k]).toLowerCase().indexOf(value) !== -1;
                }
                return false;
            });
        });
    }
}
