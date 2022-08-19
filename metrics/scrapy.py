import json  # json 包，用于读取解析，生成json格式的文件内容
import re  # 正则表达式

import requests  # 请求包  用于发起网络请求
from bs4 import BeautifulSoup  # 解析页面内容帮助包


def get_data(url):
    """
    获取数据
    :param url: 请求网址
    :return:返回请求的页面内容
    """
    # 请求头，模拟浏览器，否则请求会返回418
    header = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) '
                      'Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363'}
    resp = requests.get(url=url, headers=header)  # 发送请求
    if resp.status_code == 200:
        # 如果返回成功，则返回内容
        return resp.text
    else:
        # 否则，打印错误状态码，并返回空
        print('返回状态码：', resp.status_code)
        return ''


def parse_data(html: str = None):
    """
    解析数据
    :param html:
    :return:返回书籍信息列表
    """
    bs = BeautifulSoup(html, features='html.parser')  # 转换页面内容为BeautifulSoup对象
    kind = bs.find(name='span', attrs={'class': "now"})
    book_kind = kind.text.strip()
    ul = bs.find(name='ul', attrs={'class': 'chart-dashed-list'})  # 获取列表的父级内容

    lis = ul.find_all('li', attrs={'class': re.compile('^media clearfix')})  # 获取图书列表

    books = []  # 定义图书列表
    for li in lis:
        # 循环遍历列表
        strong_num = li.find(name='strong', attrs={'class': 'fleft green-num-box'})  # 获取书籍排名标签
        book_num = strong_num.text  # 编号
        h2_a = li.find(name='a', attrs={'class': 'fleft'})  # 获取书名标签
        book_name = h2_a.text  # 获取书名
        p_info = li.find(name='p', attrs={'class': "subject-abstract color-gray"})  # 书籍说明段落标签

        book_info_str = p_info.text.strip()  # 获取书籍说明，并 去前后空格

        books.append(
            {'book_num': book_num, 'book_name': book_name, 'book_info': book_info_str, 'book_kind': book_kind}
        )  # 将内容添加到列表

    return books


def save_data(res_list):
    """
    保存数据
    :param res_list: 保存的内容文件
    :return:
    """
    with open('books.json', 'w', encoding='utf-8') as f:
        res_list_json = json.dumps(res_list, ensure_ascii=False)
        f.write(res_list_json)


if __name__ == '__main__':

    urlList = [
        'https://book.douban.com/chart?subcat=literary',
        'https://book.douban.com/chart?subcat=novel',
        'https://book.douban.com/chart?subcat=history',
        'https://book.douban.com/chart?subcat=social',
        'https://book.douban.com/chart?subcat=comics',
        'https://book.douban.com/chart?subcat=suspense_novel',
        'https://book.douban.com/chart?subcat=business',
        'https://book.douban.com/chart?subcat=drama',
    ]

    allBook = []
    for url in urlList:
        html = get_data(url=url)  # 获取数据
        books = parse_data(html)  # 解析数据
        allBook.extend(books)

    save_data(allBook)  # 保存数据
    print('done')
