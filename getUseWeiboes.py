import requests
import time


def usr_info(usr_id): 
    url = 'https://m.weibo.cn/api/container/getIndex?type=uid&value={usr_id}'.format(usr_id=usr_id) 
    resp = requests.get(url, headers="", cookies="")
    print("Status Code:",resp.status_code)
    jsondata = resp.json() 
    print(jsondata)
    uid = jsondata.get('userInfo').get('toolbar_menus')[0].get('params').get('uid') 
    fid = jsondata.get('userInfo').get('toolbar_menus')[1].get('actionlog').get('fid') 
    oid = jsondata.get('userInfo').get('toolbar_menus')[2].get('params').get('menu_list')[0].get('actionlog').get('oid') 
    cardid = jsondata.get('userInfo').get('toolbar_menus')[1].get('actionlog').get('cardid') 
    containerid = jsondata.get('tabsInfo').get('tabs')[0].get('containerid') 
    Info = {'uid': uid, 'fid': fid, 'cardid': cardid, 'containerid': containerid, 'oid': oid} 
    print(Info) 
    return Info

def mblog_list(uid, oid): 
    Mblog_list = [] 
    base_url = 'https://m.weibo.cn/api/container/getIndex?containerid={oid}&type=uid&value={uid}' 
    page_url = 'https://m.weibo.cn/api/container/getIndex?containerid={oid}&type=uid&value={uid}&page={page}' 
    url = base_url.format(oid=oid, uid=uid) 
    resp = requests.get(url, headers="", cookies="") 
    resp.encoding = 'gbk' 
    response = resp.json() 
    total = response['cardlistInfo']['total'] 
    page_num = int(int(total) / 10) + 1 
    for i in range(1, page_num + 1, 1): 
        p_url = page_url.format(oid=oid, uid=uid, page=i) 
        page_resp = requests.get(p_url, headers="", cookies="") 
        page_data = page_resp.json() 
        cards = page_data['cards'] 
        for card in cards: 
            mblog = card['mblog'] 
            created_at = mblog['created_at'] 
            id = mblog['id'] 
            text = mblog['text']   
            reposts_count = mblog['reposts_count']     
            comments_count = mblog['comments_count'] 
            attitudes_count = mblog['attitudes_count'] 
            mblog_data = {'created_at': created_at, 
                          'id': id,
                          'text': text,
                          'reposts_count': reposts_count, 
                          'comments_count': comments_count, 
                          'attitudes_count': attitudes_count} 
            Mblog_list.append(mblog_data) 
            print(' ' * 10, mblog_data) 
        time.sleep(1) 
    return Mblog_list

def get_comments(wb_id): 
    Data = [] 
    url = 'https://m.weibo.cn/api/comments/show?id={id}'.format(id=wb_id) 
    page_url = 'https://m.weibo.cn/api/comments/show?id={id}&page={page}' 
    Resp = requests.get(url, headers="", cookies="") 
    page_max_num = Resp.json()['max'] 
    for i in range(1, page_max_num, 1): 
        p_url = page_url.format(id=wb_id, page=i) 
        resp = requests.get(p_url, cookies="", headers="") 
        resp_data = resp.json()
        data = resp_data.get('data') 
        for d in data: 
            review_id = d['id'] 
            like_counts = d['like_counts'] 
            source = d['source'] 
            username = d['user']['screen_name'] 
            image = d['user']['profile_image_url'] 
            verified = d['user']['verified'] 
            verified_type = d['user']['verified_type']     
            profile_url = d['user']['profile_url'] 
            comment = d['text'] 
        time.sleep(1)

usr_info("5129413011")

