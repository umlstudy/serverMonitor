import * as React from 'react';

// tslint:disable-next-line:interface-name
interface CheckBoxProps {
    checked:boolean;
    msg:string;
    checkBoxClick:(checked:boolean)=>void;
}

export default class CheckBox extends React.Component<CheckBoxProps> {

    // 아이콘 사이트
    // expo vector icons expo.github.io/vector-icons/

    public render() {
        const checked = this.props.checked;
        const handleCheck = this.handleCheck.bind(this);
        
        console.log("CheckBox - render checked " + checked);

        return (
            <span>
                <input type="checkbox" onChange={handleCheck} checked={checked}/>
                { !!this.props.msg ? this.props.msg : '' }
            </span>
        );
    }

    private handleCheck() {
        this.props.checkBoxClick(!this.props.checked);
    }
}