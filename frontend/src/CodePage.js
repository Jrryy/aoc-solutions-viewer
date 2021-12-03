import { useEffect, useState } from 'react';
import { CopyBlock, hybrid } from 'react-code-blocks';
import Spinner from './Spinner';

function CodePage(props) {
    const [code, setCode] = useState(null);
    const [loading, setLoading] = useState(true);

    const [executing, setExecuting] = useState(false);
    const [results, setResults] = useState(null);

    useEffect(() => {
        setResults(null);
        setLoading(true);
        fetch('AoC-2021/' + props.day)
            .then(resp => resp.text())
            .then(codeText => setCode(codeText))
            .then(() => setLoading(false));
    }, [props.day]);

    const processInput = event => {
        setExecuting(true);
        setResults(null);
        const formData = new FormData();
        formData.append('File', event.target.files[0]);
        fetch('exec/' + props.day, {
            method: 'POST',
            body: formData,
        })
            .then(resp => resp.text())
            .then(resultsText => setResults(resultsText))
            .then(() => setExecuting(false));
    };

    return loading ? (
        <Spinner />
    ) : (
        <div className="flex flex-row">
            <div className="text-sm">
                <CopyBlock
                    text={code}
                    language="python"
                    theme={hybrid}
                    codeBlock
                />
            </div>
            <div className="pl-10">
                <input
                    type="file"
                    placeholder="No input"
                    accept="text/plain"
                    onChange={processInput}
                />
                {!results && executing && <Spinner />}
                {results &&
                    results.split('\n').map(result => {
                        return <p>{result}</p>;
                    })}
            </div>
        </div>
    );
}

export default CodePage;
