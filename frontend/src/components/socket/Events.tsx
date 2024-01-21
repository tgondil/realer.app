import React from 'react';

// @ts-ignore
export function Events({ events }) {
    return (
        <ul>
            {
                events.map((event: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | Iterable<React.ReactNode> | React.ReactPortal | null | undefined, index: React.Key | null | undefined) =>
                    <li key={ index }>{ event }</li>
                )
            }
        </ul>
    );
}