import { useEffect, useState } from 'react';
import { CopyBlock, hybrid } from 'react-code-blocks';
import Spinner from './Spinner';

function CodePage(day) {
    const [code, setCode] = useState({});
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        setLoading(true);
        fetch('AoC-2021' + window.location.pathname)
            .then(resp => resp.text())
            .then(codeText => setCode(codeText.toString()))
            .then(() => new Promise(resolve => setTimeout(resolve, 10000)));
        setLoading(false);
    }, [day]);

    return loading ? (
        <Spinner />
    ) : (
        <div className="text-sm">
            <CopyBlock text={code} language="python" theme={hybrid} codeBlock />
        </div>
    );
}

export default CodePage;
