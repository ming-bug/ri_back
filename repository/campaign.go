package repository

import (
	. "ri/driver"
	. "ri/model"
)

// QueryCampaignByEmail Query campaign by email
func QueryCampaignByEmail(email string, current int, pagesize int) ([]Campaign, error) {
	campaigns := make([]Campaign, 0)
	//err := Engine.Cols("CAMPAIGNUIN", "CAMPAIGNNAME", "STATUS", "CREATOR", "CREATETIME", "LASTMODIFYTIME").Table("CAMPAIGN").Where("CREATOR = ?", email).Desc("LASTMODIFYTIME").Limit(pagesize*current, pagesize*(current-1)).Find(&campaigns)
	err := Engine.SQL("SELECT CAMPAIGNUIN, CAMPAIGNNAME, STATUS, INTIMATECAMPAIGNTYPE, CREATOR, TO_CHAR(CREATETIME/(1000*60*60*24)+ TO_DATE('1970/01/01 08:00:00', 'YYYY/MM/DD HH24:MI:SS'), 'YYYY/MM/DD HH24:MI:SS') CREATETIME, TO_CHAR(LASTMODIFYTIME/(1000*60*60*24)+ TO_DATE('1970/01/01 08:00:00', 'YYYY/MM/DD HH24:MI:SS'), 'YYYY/MM/DD HH24:MI:SS') LASTMODIFYTIME, TO_CHAR(ACTUALSTARTTIME/(1000*60*60*24)+ TO_DATE('1970/01/01 08:00:00', 'YYYY/MM/DD HH24:MI:SS'), 'YYYY/MM/DD HH24:MI:SS') ACTUALSTARTTIME, TO_CHAR(ACTUALENDTIME/(1000*60*60*24)+ TO_DATE('1970/01/01 08:00:00', 'YYYY/MM/DD HH24:MI:SS'), 'YYYY/MM/DD HH24:MI:SS') ACTUALENDTIME FROM (SELECT CAMPAIGNUIN, CAMPAIGNNAME, STATUS, INTIMATECAMPAIGNTYPE, CREATOR, CREATETIME, LASTMODIFYTIME, ACTUALSTARTTIME, ACTUALENDTIME, ROWNUM RN FROM (SELECT CAMPAIGNUIN, CAMPAIGNNAME, STATUS, INTIMATECAMPAIGNTYPE, CREATOR, CREATETIME, LASTMODIFYTIME, ACTUALSTARTTIME, ACTUALENDTIME FROM CAMPAIGN WHERE (CREATOR = :1) ORDER BY LASTMODIFYTIME DESC) at WHERE (ROWNUM <= :2)) aat WHERE (RN > :3)", email, current*pagesize, (current-1)*pagesize).Find(&campaigns)
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}
